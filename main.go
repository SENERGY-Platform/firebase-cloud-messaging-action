/*
 * Copyright 2023 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"log"
	"os"
)

func main() {
	credentials, ok := os.LookupEnv("INPUT_CREDENTIALS")
	if !ok {
		log.Fatal("expected env INPUT_CREDENTIALS to be set")
	}
	credentialMap := map[string]string{}
	err := json.Unmarshal([]byte(credentials), &credentialMap)
	if err != nil {
		log.Fatal("env INPUT_CREDENTIALS is not json")
	}

	projectId, ok := credentialMap["project_id"]
	if !ok {
		log.Fatal("expected env INPUT_CREDENTIALS to contain key project_id")
	}

	messageRaw, ok := os.LookupEnv("INPUT_MESSAGE")
	if !ok {
		log.Fatal("expected env INPUT_MESSAGE to be set")
	}
	message := messaging.Message{}
	err = json.Unmarshal([]byte(messageRaw), &message)
	if err != nil {
		log.Fatal("env INPUT_MESSAGE cannot be unmarshalled into messaging.Message")
	}

	if err := os.WriteFile("credentials.json", []byte(credentials), 0440); err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = os.Remove("credentials.json")
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "credentials.json")
	if err != nil {
		log.Fatal("could not set env GOOGLE_APPLICATION_CREDENTIALS")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	firebaseApp, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: projectId,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	firebaseClient, err := firebaseApp.Messaging(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := firebaseClient.Send(ctx, &message)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(resp)
}

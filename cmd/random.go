/*
Copyright © 2021 Alex Wasik arthur.wasik@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random joke",
	Long:  `Get a random joke`,
	Run: func(cmd *cobra.Command, args []string) {
		const MAX_JOKE_ID int = 377
		data, err := ioutil.ReadFile("./files/jokes.json")
		if err != nil {
			fmt.Println(err)
		}

		var objmap []map[string]interface{}
		err = json.Unmarshal(data, &objmap)
		if err != nil {
			fmt.Println("Error:", err)
		}

		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(MAX_JOKE_ID)
		fmt.Println(objmap[randomIndex]["setup"])
		fmt.Println(objmap[randomIndex]["punchline"])
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

// Copyright 2020 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/rakyll/go-sql-driver-spanner"
)

func main() {
	ctx := context.Background()
	db, err := sql.Open("spanner", "projects/PROJECT/instances/INSTANCE/databases/DATABASE")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.QueryContext(ctx, "SELECT id, text FROM tweets WHERE likes > @likes LIMIT 10", 100)
	if err != nil {
		log.Fatal(err)
	}
	var (
		id   int64
		text string
	)
	for rows.Next() {
		if err := rows.Scan(&id, &text); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, text)
	}
	rows.Close()
}

package main

var Url string = "https://gitlab.com/api/graphql"
var Req string = `
		query last_projects($n: Int = DISPLAY_NUM) {
			projects(last:$n) {
				nodes {
					name
					description
					forksCount
				}
			}
		}
	`

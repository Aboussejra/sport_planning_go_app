build_and_launch:
	go build && ./sport_planning_go_app
develop:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
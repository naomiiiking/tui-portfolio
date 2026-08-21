package content

import "strings"

const Name = "NAOMIKING.SSH"

const Title = 
`    __        *             _                 *     
  /\ \ \__ _  ___  _ __ ___ (_)       *  _         *    
 /  \/ / _' |/ _ \| '_ ' _ \| |     /\ /(_)_ __   __ _
/ /\  / (_| | (_) | | | | | | |    / //_/ | '_ \ / _' |
\_\ \/ \__,_|\___/|_| |_| |_|_|   / __ \| | | | | (_| |
	    *					*    \/  \/|_|_| |_|\__, |
				   *					*		 |___/ `
const About = `Founding engineer for the AI Solutions team at Good Systems (formerly Goodnotes). Building AI systems for production. Not currently looking for work, but to expand my network with women in STEM and AI Engineers in London.`


type Job struct {
	Title    string
	Company  string
	Period   string
	Location string
	Bullets  []string
}
var Experience = []Job{
	{
		Title:    "AI Engineer",
		Company:  "Goodnotes",
		Period:   "2025 – present",
		Location: "London, UK",
	},
	{
		Title:    "Customer Support Engineer",
		Company:  "Goodnotes",
		Period:   "2024 – 2025",
		Location: "London, UK",
	},
	{
		Title:    "Technical Expert",
		Company:  "Apple",
		Period:   "2021 – 2024",
		Location: "Brisbane, Australia",
	},
}

type Skill struct {
	Title    string
	Examples []string
}

var Skills = []Skill{
	{
		Title: "AI Development",
		Examples: []string{
			"LLMs: Fine tuning, orchestration, deployment, vision",
			"AI Frameworks: Langraph, Langchain, AI SDK, HuggingFace",
			"RAG: Advance production hybrid vector retrieval with pgvector, FAISS",
			"MCP: Building custom remote MCP servers, RBAC tool usage through OAuth and OIDC",
			"ML (Currently learning): Sckitlearn, PyTorch",
		},
	},
	{
		Title: "Software engineering",
		Examples: []string{
			"Programming: Python, Typescript",
			"Servers: Flask, Express.js, Bun",
			"Chat: ChatSDK, Bolt",
			"Integrations: Fast API, webhooks, REST, GraphQL",
		},
	},
	{
		Title: "Infrastructure",
		Examples: []string{
			"Cloud Compute: AWS, VPS hosting, Neoclouds, Google Cloud Services",
			"Containerisation: Docker, Kubernetes, ArgoCD",
			"GitOps: Terraform, CI/CD",
		},
	},
}

type Link struct {
	Label string
	URL   string
}

var Links = []Link{
	{Label: "github", URL: "https://github.com/naomiiiking"},
	{Label: "linkedin", URL: "https://www.linkedin.com/in/naomi-king-0374891ba/"},
}

const rawPortrait = 
`......................................................................
.......................................................................
.......................................................................
.......................................................................
.......................................................................
.......................................................................
.........................--.    .......................................
......................:=**=-=-::+=:-...................................
.....................=####*#@@@@@%*. %.................................
....................=##%%%%@@@@@%#*:  .................................
...................=#*+***#@#+*%#+=:   ................................
..................:#*=-=*%%#=.#@-:-:   ................................
..................=++::+@@@@@@@@+:-    ................................
.................:=--::+%@%=+@%=..+#..  ...............................
.................=:.:.-+%@+-*#+:. .:.  :...............................
................:-.:..:=%@@@@#-.  ..   :...............................
...................:..::*@##@%+..-.  . :...............................
..................:-.....   .++:..     ................................
.................:=+.:... .  .::.   .. ................................
................:=++--...     ..    ..  .+*#*-:........................
...........-=*##@##=-=:.::    ...   ..:=*@@@@@@*:......................
........=@@@@@@@@@@::-:.-##=**#%##.:::%@@@@@@@%*=:.....................
.......*@@@@@@@@@@.-=:::+@@@@%#@@@=--#+%@@@@@%#+-:.....................
......-@@@@@@@@@@:-*@@@@@@@@%**=*@=--=@@@@@@@%*=-:.....................
......*@@@@@@@@@*:-#@@@@@@@@@@-.:@=::-#@@@@@%#*+-:.....................
.....:%@@@@@@@@%===#@@@@@@@@@#*%%*.::::%@@#*##*=-:.....................
.....-@@@@@@@@@+==.=@@@@@@@@@#:     :-:.*+-:=#*=-:.....................
.....-@@@@@@@%*+=-:=@@@@@@@%-+.  ..-:---:-..:*+=-:.....................
.....=@@@@@@@#-+===:%@@@@@%-+@=   --::-::.. .==--::....................
.....+@@@@@@@*-===::#@@@#+:#@@+:*:    .::. ..-==--:....................
....:+@@@@@@%=++--@@@@%=.:%@#--..-    .:.   .:==--:....................
....:*@@@@@@#==#@@@@@*:..-+=-:.              .==--:....................
....:*@@@@@@@@@@@@@#=..                       :==--:...................
....:*@@@@@@@@@@@%+:.                       . .-=--:...................
....:*@@@@@@@@@%*:.                         .. .-----:.................
....:+@@@@@@%#=:..               ....       .. .-=++=-.................
.....-@@@%#+-...               .::::::     ...  :++**=:................
.....:*%*-....        .::. ..  :--:::.:.   ...  :*##*+:................
.............       .:::::. ...---::.....  .... :*##*=:................`

var Portrait = func() string {
	lines := strings.Split(rawPortrait, "\n")
	return strings.Join(lines, "\n")
}()

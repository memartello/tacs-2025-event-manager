# tacs-2025-event-manager


# INFRASTRUCTURE
---
A docker compose that get docker go backend and react frontend and setup a proxy server (nginx) to have backend and front on the same domain.


Maybe an script that run the backend with the front.

# GOLANG BACK
---
[] Code quality tools?
Workspace using different modules.
for adding a new module go work use "go work use ./newmodule"



# REACT WEB
---
Vite Tempalte React-swc-ts
React-query
Axios
Shadcn as library with Tailwind css

Setup (Linter and code quality)




Running the docker entire
---
docker compose -f infra/docker-compose.yml up --build -d
This will expose backend in PORT 8080 and frontend in 3000.
Both under the same proxy server and with this golang api can be running with cors disabled.
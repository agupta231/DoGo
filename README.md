# DoGo

## Context
A **Do**cker + **Go** playground. Really just a repository to learn the new Google language and the
"deployment technology of the future."

## Disclaimer
This is a goofing around / learning a new language type of project, and most of the code doesn't
conform to "good programming practices". This project also has no plans to be maintained. I urge
you to use your common sense if you decide to reference parts of this project.

## Execution Directions 
Clone the project, and simply run `docker-compose run client`, and in theory, the whole system
should start right up. 

In this case, `docker-compose up` doesn't work because client requires an interactive shell. Docker
Compose is really optimized for servers, so you have to deal with a little bit of jank if you are
creating CLI applications.

# DoGo

## Context
A **Do**cker + **Go** playground. Really just a repository to learn the new Google language and the
"deployment technology of the future."

### Primary Learning Objectives
As this is really a learning experience more than anything, the following is a list of things I
wanted to mess around with in this project. If any of these topics are relevant to what you are
doing, you may want to check out the source.

* Use Go to create a REST API server
* Use Go to create a REST API client
* Use Docker to isolate the whole setup
* Use Docker Compose to spin up different services and have everything work together magically
* Use Docker Compose to create a private network in which the different services can talk to each
  other, but no outside programs can access
* Use multi-stage builds to reduce image sizes

## Disclaimer
This is a goofing around / learning a new language type of project, and most of the code doesn't
conform to "good programming practices". This project also has no plans to be maintained. I urge
you to use your common sense if you decide to reference parts of this project.

## Execution Directions 
### Development
Clone the project, and simply run `docker-compose run client`, and in theory, the whole system
should start right up. 

### Production
The production variant of the images have the same functionality, but a much smaller footprint.
Clone the project, and run `docker-compose -f docker-compose-prod.yml run client`. The main reason
why you should use this full time is because it copies the source code into the docker image in
order to compile. Thus, if an edit is made to the code, the whole project needs to be rebuilt.

In contrast, the development mode rebuilds the project in runtime, allowing it to reuse the old (but
bigger) images. Also, as the environment has already been created, the build times are much faster, 
which is a godsend when working on the project.


**Note:** In this case, `docker-compose up` doesn't work because client requires an interactive 
shell. Docker Compose is really optimized for servers, so you have to deal with a little bit of 
jank if you are creating CLI applications.

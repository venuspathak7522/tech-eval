# Skills Evaluation Test
The goal of this process is evaluate SRE candidates overall skills on:
1. General RESTful service architecture
2. Coding and debugging
3. Usage of operation tools

Those abilities will be evaluated based on a series of problems that need to be solved
based on a RESTful service skeleton intentionally written with faults to create a scenario
where the candidate may be able to diagnose and solve any blocking condition.

Further on, additional challenges are presented assuming the basic service is working properly.
Those challenges give to the candidate the opportunity to explore solutions on his/her own,
not necessarily tied to a predetermined answer.

## Evaluation

Welcome to Qlik SRE Technical Evaluation!

First of all, thank you for your time and your interested in joining our team. We tried to elaborate
a process of technical evaluation that gives you freedom of thinking at the same time that
allows us to understand better your technical skills.

Before you begin, please ensure you read over the following guidelines:

1. You received versions of the same service in different programming languages. Choose ONE!
Preferably the one that you are more comfortable with.
2. Create a public repo in Github and upload the source code with the version you chose.
3. For each task, try to put everything together in one single Pull Request. This will make it easier for us
to review what you have done.
4. Feel free to create several Docker images in you desire. If you choose so, please publish each image with an unique tag.

**Remember!**
* There are no expectations that you will complete ALL of them.
* The idea here gauge your knowledge with the tools we, as SRE, use on a regular basis.
* Try to complete tasks that you are more comfortable first. Leave the ones you are less familiar for

Qlik's Culture and Talent team will contact you to setup a virtual meeting
where you will be required to share your screen and participate in a code review of your PRs with
members of Qlik's SRE Team. As with any code review, we will be inquiring about design choices,
coding styles, challenges and potential enhancements.

## The scenario

We have this very simple RESTful service. It is used internally by the SRE team to create
a library of Palindromes. Someone used to have fun writting those once in a while and it
kind of became a tradition. Unfortunately, whoever started this project never finished it
completely and it got merged with a series of problems.

The project has 5 endpoints:

| METHOD | Endpoint              | Description                            |
|--------|-----------------------|----------------------------------------|
| GET    | /api/v1/messages      | Returns all messages                   |
| GET    | /api/v1/messages/[ID] | Returns a specific record              |
| POST   | /api/v1/messages      | Creates a new record                   |
| DELETE | /api/v1/messages/[ID] | Deletes an existing record             |
| GET    | /api/v1/health        | Displays service overall health status |

## Tasks

The following tasks have been assigned to you to complete on this project:

1. **Fix the project (EASY):** Make sure it is working and that the `messages` endpoints do what they are suppose to.
2. **Dockerize it (EASY):** The project has a `Dockerfile` that should be able to handle this service. Try to
figure out what is wrong and make sure the container is accessible
3. **Where is the palindrome? (MEDIUM):** This service is NOT doing what is supposed to. It records several
messages but it doesn't validate if they are Palindromes. Add a boolean flag to the Message
models where it says whether or not it IS a palindrome. Hint: create a method that validate if the
word IS a palindrome before it is saved.
4. **Where is the persistence? (HARD):** The service is storing all the records in memory. This is bad because
we want to have multiple instances running. Convert the data layer to store the records in a MongoDB database.

**BONUS TASK**

1. **We need more (MEDIUM):** The `health` endpoint is supposed to show us what is going on with the service.
Things like, CPU, memory, disk usage, # of requests, # of errors (You got extra points is you use [prometheus libraries](https://prometheus.io/docs/instrumenting/clientlibs/))

2. **Concurrency is the key(HARD):** If you completed task #4, it is time to make this service
cloud available. You have the choice to implement this using Docker Swarm or Kubernetes.
   1. Bring up a MongoDB database
   2. Run 4 instances of the service, all connecting to the same MongoDB database.
   3. Create a service mapping all 4 instances.

**Note:** If you choose to implement Bonus Task #2 using Kubernetes, make sure you have a Helm chart or a single
YAML file with the manifests.

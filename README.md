## SRE ASSESMENT - Muhamad Deni Afendi

1. I use boiler plate in this repository reference https://github.com/Papagoat/go-bootstrap-boilerplate/tree/master/go-bootstrap/4

2. CI/CD 
- a. build - use go build
![image](https://user-images.githubusercontent.com/80587939/172310577-f100c362-ba79-402e-8024-28006d2532d1.png)
- b. build container - docker build with dockerfile
![image](https://user-images.githubusercontent.com/80587939/172310600-e16ab941-28c4-439e-8c5f-6b9fe9537211.png)
- c. store ro docker hub
![image](https://user-images.githubusercontent.com/80587939/172310105-5563ba7b-aa17-41f6-b64b-037bf4ed6efc.png)
- d. deploy to kubernetes (GKE)
![image](https://user-images.githubusercontent.com/80587939/172310813-785023e5-475c-4e59-b52f-685bf53e4407.png)
![image](https://user-images.githubusercontent.com/80587939/172310920-075aeccd-1945-4993-871d-e267366422c0.png)

3. service http://belajarsre.sytes.net/
![image](https://user-images.githubusercontent.com/80587939/172311106-77cfa32f-b9a9-4320-8a31-5cc0c29976eb.png)

4. Tech Stack 
- For programming languange: [`golang`]
- Repository: [`Github`]
- Docker registrry: docker hub
- CI/CD: github action path gosimpleapp/.github/workflows/
  note: Github action trigger by push and pull request to main branch
- Container Orchertration: [`GKE (Google Kubernetes Engine)`]
- Service & ingress: service loadbalancer google



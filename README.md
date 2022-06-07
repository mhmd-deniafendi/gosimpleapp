## SRE ASSESMENT - Muhamad Deni Afendi

1. I use boiler plate in this repository reference https://github.com/Papagoat/go-bootstrap-boilerplate/tree/master/go-bootstrap/4

2. CI/CD 
- a. build - use go build
![image](https://user-images.githubusercontent.com/80587939/172310577-f100c362-ba79-402e-8024-28006d2532d1.png)

- b. build container - docker build with dockerfile
![image](https://user-images.githubusercontent.com/80587939/172314214-f10c10b5-2c0f-42a4-9939-0c158facd7b8.png)

- c. Store container image to Docker Registry
![image](https://user-images.githubusercontent.com/80587939/172310105-5563ba7b-aa17-41f6-b64b-037bf4ed6efc.png)

- d. Deploy to kubernetes cluster (GKE)
![image](https://user-images.githubusercontent.com/80587939/172313602-d2019691-ae3a-4f0f-b08f-fd8d9ab1a3dc.png)
![image](https://user-images.githubusercontent.com/80587939/172313746-2121abfe-ac36-4fbf-8c19-062275542288.png)
![image](https://user-images.githubusercontent.com/80587939/172314015-572f7d81-c556-4d1f-b056-22e402ce0d94.png)

3. service http://belajarsre.sytes.net/ = i user https://my.noip.com/
![image](https://user-images.githubusercontent.com/80587939/172311106-77cfa32f-b9a9-4320-8a31-5cc0c29976eb.png)

4. Tech Stack 
- For programming languange: golang
- Repository: Github
- Docker registrry: docker hub
- CI/CD: github action path gosimpleapp/.github/workflows/
  note: Github action trigger by push and pull request to main branch
- Container Orchertration: GKE (Google Kubernetes Engine)
- Service & ingress: service loadbalancer google



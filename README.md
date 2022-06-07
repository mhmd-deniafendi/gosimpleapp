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

3. service http://belajarsre.sytes.net/ = i use https://my.noip.com/
![image](https://user-images.githubusercontent.com/80587939/172311106-77cfa32f-b9a9-4320-8a31-5cc0c29976eb.png)

4. Tech Stack 
- For programming languange: golang
- Repository: Github
- Docker registrry: docker hub
- CI/CD: github action path gosimpleapp/.github/workflows/
  note: Github action trigger by push and pull request to main branch
  
  ![image](https://user-images.githubusercontent.com/80587939/172317552-7d4b16a8-9659-4db7-9e79-33e949a472e8.png)
  
- Container Orchertration: GKE (Google Kubernetes Engine)
- Service & ingress: service loadbalancer google


## Explanation
1. I use this https://github.com/Papagoat/go-bootstrap-boilerplate/tree/master/go-bootstrap/4 repository as a reference 
2. Create CI/CD
- I just build, build with this command with dockerfile :
```
$ go build -o PINTU .
```

![image](https://user-images.githubusercontent.com/80587939/172316703-bcb18793-c41e-4e4b-9e9d-551db67c6eb7.png)


- Building container image, handle by module github action:

![image](https://user-images.githubusercontent.com/80587939/172315777-97cc8728-39d8-4c73-84d1-66de60afa88a.png)

- Store to Docker hub as a docker registry

- Deploy to kubernetes cluster: i use googke kubernetes engine
![image](https://user-images.githubusercontent.com/80587939/172316178-69a7148c-01cf-469a-8b17-d4ffd922a48b.png)

3. DNS Domain: i use https://my.noip.com/
![image](https://user-images.githubusercontent.com/80587939/172316331-e887c02b-378d-4d42-841e-2c2dbd423da0.png)

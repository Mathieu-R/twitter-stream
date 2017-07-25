# twitter-stream
Simple tweeter tracker 
> Note : this is not a big project, it's a tiny one for me to learn go

### Utilization :  

#### Grab the project
```
go get github.com/Mathieu-R/twitter-stream
```

#### Credentials
You need to create a twitter apps to get your credentials : https://apps.twitter.com/

Then, you should set the following environment variables

- CONSUMER_KEY
- CONSUMER_SECRET


#### Launch the app 
```
$ twitter-stream -track <keyword>
```
```
$ twitter-stream -track <keyword1, keyword2,...>
```

#### Docker Image 
```
docker pull matiuso/twitter-stream
```
```
docker run matiuso/twitter-stream -e CONSUMER_KEY=<key> -e CONSUMER_SECRET=<secret> -track <keyword>
```

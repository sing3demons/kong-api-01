# kong-api-01

```start go-service
cd go-server
go run main.go
```

```start node-js
cd node-js
npm install
npm start
```

```start kong
docker-compose up -d
```

<a href="http://localhost:1337"> go to http://localhost:1337 </a>

<ul>
<li>KONG ADMIN URL : http://kong-gateway:8001</li>
<li>SERVICE :: <button>ADD NEW SERVICE</button> => <h6>Url</h6> => http://host.docker.internal:8080/ => Routes <button>+ ADD ROUTE</button> +> Paths=/,Methods=GET</li>
<li>SERVICE :: <button>ADD NEW SERVICE</button> => <h6>Url</h6> => http://host.docker.internal:3000/ => Routes <button>+ ADD ROUTE</button> +> Paths=/,Methods=GET</li>
</ul>

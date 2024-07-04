docker build -t <your-controller-image> .
docker push <your-controller-image>


```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: shipton-controller-manager
  namespace: shipton-system
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: controller-manager
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
        - name: manager
          image: <your-controller-image>
          command:
            - /manager
          args:
            - --metrics-addr=:8080
            - --enable-leader-election
          ports:
            - containerPort: 8080
              name: metrics
```
kubectl apply -f deploy.yaml


```
apiVersion: shipton.yourdomain.com/v1alpha1
kind: ShiptonBuild
metadata:
  name: my-shipton-build
spec:
  image: <your-docker-registry>/<your-image-name>:latest
  dockerfile: /path/to/Dockerfile
  context: /path/to/context
```

kubectl apply -f my-shipton-build.yaml

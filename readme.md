# Deployer
## Quick Start
### Debug mode (Development mode)
```shell
project-path> go run main.go
```

### Release mode (Production mode)
```shell
proejct-path> go run main.go [full-path-to-project]
```

connect with localhost:8080

http://localhost:8080

## Install and start with system unit

1. Install deployer project
```shell
project-path> go install
```

2. Edit **deployer.service** file

Original
```text
ExecStart=/GOBIN/deployer /PROJECT-PATH
```

Example
```text
ExecStart=/home/pi/go/bin/deployer /home/pi/deployer
```
 
3. Copy system unit
```shell
project-path> sudo cp ./deployer.service /etc/systemd/system/deployer.serive
```

4. Start system unit
```shell
> sudo systemctl start deployer.service
```

5. (Optional) Restart unit on reboot
```shell
> sudo systemctl enable deployer.service
```

6. connect with localhost:8080


http://localhost:8080
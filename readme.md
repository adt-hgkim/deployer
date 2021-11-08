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

## Install and start with system unit

1. Install deployer project
```shell
project-path> go install
```

2. Edit **deployer.service** file
```text
ExecStart=/GO-BIN-PATH/deployer /PROJECT-PATH
```
```text
ExecStart=/home/pi/go/bin/deployer /home/pi/deployer
```
 
3. Copy system unit
```shell
> sudo cp ./deployer.service /etc/systemd/system/deployer.serive
```

4. Start system unit
```shell
> sudo systemctl start deployer.service
```

5. (Optional) Restart unit on reboot
```shell
> sudo systemctl enable deployer.service
```

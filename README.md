# Sensu_mattermost
sensu go handler for mattermost, independent on OS

#Get dep
>export GO111MODULE="on" 
>go get github.com/urfave/cli/v2
>go mod init
#HowToBuild
>go build -ldflags "-w -s" -o mattermost


#HowTOUSE
>mattermost 
>mattermost --webhook|-w https://mattermost.mycompany.com/hooks/balabalababalababa



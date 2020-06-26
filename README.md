# Sensu_mattermost
sensu go handler for mattermost, independent on OS

#Get dep
>export GO111MODULE="on" 
>go get github.com/urfave/cli/v2
>go mod init

#HowTo
##Build Binary locally
>go build -ldflags "-w -s -X main.version=${version} -X main.author=WenZhou"  

##Build Binary from docker (if no Go env set locally) and upload to artifactory  
>docker build --force-rm=true --build-arg version=0.1.1 --build-arg component=mattermost .


#HowTOUSE
>mattermost --help
##interactive mode
>mattermost --webhook https://mattermost.mycompany.com/hooks/balabalababalababa
##detached mode
>cat info.sh
export topic="Party"
export time="tomorrow 2020-20-20 20:20 P.M"
export announcement="we <3 :beer: please join us"
>source info.sh
>mattermost an -d


#Usage(example)
##manual
>mattermost an [--webhook https://mattermost.mycompany.com/hooks/abcdabcdabcdabcd]
Enter Announcement Topic: (e.g Jenkins Upgrade, LDAP change, etc) Update Jenkins:   
Enter Announcement detail(stop by 'Wen'):  
when (e.g yyyy-mm-dd@hh:mm or near future etc): on 2020-20-20 20:20 P.M   
    Detail line: We will schedule jenkins upgrade from version 2.189 to 2.190.3 :D  
    Detail line: Users can expect 1hr-ish downtime and be informed once the job is done   
    Detail line: For any question/concern, please reach out to me <3  
    Detail line: WEN   
##jenkins   
pass topic, time, announcement from Jenkins job as string, string, multiline string from Web GUI  
>mattermost an [--webhook https://mattermost.mycompany.com/hooks/abcdabcdabcdabcd -d]  
##Sensu handler
>mattermost al [--webhook https://mattermost.mycompany.com/hooks/efdhefdhefgh]   

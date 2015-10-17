export GOPATH=$(pwd)

alias cnn='ssh -i ~/.ssh/was.pem.txt ubuntu@50.112.163.250'
alias restart="sudo lsof -i:80 | tail -1 | awk '{print $$2}' | sudo xargs kill -9;  sudo nohup ./bin/wx &"


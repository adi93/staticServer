ecryptfs-mount-private
scp ~/Private/notes/html aditya@51.158.118.165:~/
scp ~/Private/notes/js aditya@51.158.118.165:~/


# cd ~/staticServer
# rm -r website
# git pull
# mkdir website
# cd website
# ln -s html ~/html
# ln -s js ~/js
# ~/replaceJs.sh
# cd ~/staticServer
# make
# ./staticServer


## replaceJs.sh
# cd ~/html
# find . -name "*.html"

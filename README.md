I created this project to deploy my static website.

I am pretty happy with vimwiki, but I need some sort of access control on it. I was about to go with hugo, but that would be an effort for me to find a good markdown to wiki converter. And it still wouldn't solve the access control function.

So I am going to use golang's file server. I will specify in a configuration file what directories, like 'work', I want to include/exclude.

gamedayweb
==========

### Build and run

Simple go webapp that utilizes the gamedayapi. For now, just displays the teams for 2014. More to come.

    git clone git@github.com:ecopony/gamedayweb.git
    cd gamedayweb
    go build
    ./gamedayweb

Go to http://localhost:3000/index.html

### Or do the docker thing

    cd gamedayweb
    docker build -t gamedayweb .
    docker run -p 3000:3000 gamedayweb

If you're using boot2docker, determine the IP 

    boot2docker ip

Then hit that IP on port 3000. For example: http://192.168.59.103:3000/index.html
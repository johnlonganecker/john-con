# Make your own Container in Go

### Code follows this talk by Liz Rice
<iframe width="560" height="315" src="https://www.youtube.com/embed/HPuvDm8IC-4" frameborder="0" allowfullscreen></iframe>

### Talk was based on this article
[Build Your Own Container Using Less than 100 lines of Go](https://www.infoq.com/articles/build-a-container-golang)

Hardest part was getting a rootfs (which I included in this repo). The rootfs for this example came from [Tiny Core Linux](http://tinycorelinux.net/downloads.html).

This code needs to be compiled and run from a linux. I included a Vagrantfile which should make your life a lot easier.

### Setup

```
vagrant up
vagrant ssh

sudo apt-get install golang

cd /vagrant

# run as root
sudo su

# list directories from inside our container
go run main.go run /bin/sh

# run a shell (/bin/sh) from inside our container
go run main.go run /bin/sh
```

### How to create the rootfs from a ISO file

either mount the iso or use 7z to unpack the iso

Downloaded this iso [TinyCore Core](http://tinycorelinux.net/7.x/x86/release/Core-current.iso)

**mount iso**
```
mkdir iso
mount Core-current.iso iso
```

OR

**unpack iso with 7z**
```
sudo apt-get install p7zip-full

mkdir iso
mv Core-current.iso iso
cd iso
7z x Core-current.iso
```

Then create our rootfs

```
cd boot
gzip -dk core.gz
cpio -i < core
```

you will see some operations not permitted, but this is enough files for our project

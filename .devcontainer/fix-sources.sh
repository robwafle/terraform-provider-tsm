mkdir ~/solution
cd ~/solution/

cat << EOF > ~/solution/sources.list
deb http://archive.ubuntu.com/ubuntu/ $(lsb_release -cs) main restricted universe multiverse
deb-src http://archive.ubuntu.com/ubuntu/ $(lsb_release -cs) main restricted universe multiverse

#deb http://archive.ubuntu.com/ubuntu/ $(lsb_release -cs)-updates main restricted universe multiverse
#deb-src http://archive.ubuntu.com/ubuntu/ $(lsb_release -cs)-updates main restricted universe multiverse

#deb http://archive.ubuntu.com/ubuntu/ $(lsb_release -cs)-security main restricted universe multiverse
#deb-src http://archive.ubuntu.com/ubuntu/ $(lsb_release -cs)-security main restricted universe multiverse

#deb http://archive.ubuntu.com/ubuntu/ $(lsb_release -cs)-backports main restricted universe multiverse
#deb-src http://archive.ubuntu.com/ubuntu/ $(lsb_release -cs)-backports main restricted universe multiverse

deb http://archive.canonical.com/ubuntu $(lsb_release -cs) partner
deb-src http://archive.canonical.com/ubuntu $(lsb_release -cs) partner
EOF

rm /etc/apt/sources.list
cp ~/solution/sources.list /etc/apt/sources.list

# mv /etc/apt/sources.list.d/* ~/solution

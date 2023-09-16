#!/bin/bash


date_time=$(date +"%a %b %d %H:%M:%S %Z %Y")


mkdir "$1 at $date_time"


mkdir "$1 at $date_time/about_me"


mkdir "$1 at $date_time/about_me/personal"
mkdir "$1 at $date_time/about_me/professional"



echo "https://www.facebook.com/$2" > "$1 at $date_time/about_me/personal/facebook.txt"


echo "https://www.linkedin.com/in/$3" > "$1 at $date_time/about_me/professional/linkedin.txt"


mkdir "$1 at $date_time/my_friends"

mkdir "$1 at $date_time/my_system_info"



curl -s https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt >> "$1 at $date_time/my_friends/list_of_my_friends.txt"

echo "Nama user: $(whoami)" >> "$1 at $date_time/my_system_info/about_this_laptop.txt"
echo "$(uname -a)" >> "$1 at $date_time/my_system_info/about_this_laptop.txt"

ping -c 3 google.com >> "$1 at $date_time/my_system_info/internet_connection.txt"

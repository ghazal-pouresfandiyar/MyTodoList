#! /bin/bash

case $1 in
	add)
		case $2 in
			-t | --title)
				if [[ -z $3 ]]; then
					echo "Option -t|--title Needs a Parameter"
					exit 1
				fi
			;;
		esac
		case $4 in
			-p | --priority)
				if [ -z $5 ]; then
					priority="L"
				elif [ $5 == 'M' ] || [ $5 == 'L' ] || [ $5 == 'H' ]; then
					priority=$5
				else
					echo "Option -p|--priority Only Accept L|M|H"
					exit 1
				fi
			;;
			*)
				priority="L"
			;;
		esac
		echo "0,$priority,\"$3\"" >>  tasks.csv
	;;
	clear)
		> tasks.csv
	;;
	list)
		cat tasks.csv | awk -F"," '{print NR " | " $1 " | " $2 " | " $3}'
	;;
	find)
		cat tasks.csv | grep $2 tasks.csv | awk -F"," '{print NR " | " $1 " | " $2 " | " $3}'
	;;
	done)
		sed -i -e "$2""s/0/1/" tasks.csv
	;;
	*)
		echo Command Not Supported!
	;;
esac

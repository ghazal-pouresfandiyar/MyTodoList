#! /bin/bash

case $1 in
	add)
		case $2 in
			-t | --title)
				if [[ -z $3 ]]; then
					echo "Option -t|--title Needs a Parameter"
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
				fi
			;;
			*)
				priority="L"
			;;
		esac
		if !([ -z $2 ]); then
			echo "0,$priority,$3" >>  tasks.csv
		fi
	;;
esac

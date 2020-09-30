#!/bin/sh

echo 
echo "----------------------------------------------------------------------------------"
echo "Thaiwater30 project requires postgis, postgis_topology and address_standardizer."
echo ""
echo "These PostgreSQL extensions require database administration privileges to install"
echo ",therefore it can not be install using this installation script."
echo ""
echo "Make sure that you ask your database administator to install these extensions on "
echo "the database before continue. (see: http://postgis.net/install/ for more details)"
echo "----------------------------------------------------------------------------------"
if [ ${FLAG_INPUT_USEDEFAULT} -ne 1 ]; then
	pause
fi

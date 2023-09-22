rm -rf code
mkdir code
cd code
git clone https://github.com/devfans/Zion.git
cd Zion && git checkout develop
make zion-local
cp geth ../../
cd ..
git clone https://github.com/siovanus/distribute-check.git
cd distribute-check
go build -o distribute-check
cp distribute-check ../../
cd ..
git clone https://github.com/siovanus/clean_db.git
cd clean_db
go build -o clean_db
cp clean_db ../../
cd ..
git clone https://github.com/devfans/wizard.git
cd wizard
go build -o wizard
cp wizard /usr/local/bin
cd ../../
go build -o zion-test
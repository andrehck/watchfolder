# WatchFolder

- Todos os arquivos inseridos na path assets/src são movidos para assets/dst
- Para executar go run spybond.go
- Adiciona nas variaveis de ambiente o path de origem e destino, com as seguintes key's: PATHSRC e PATHDST, SITEENV, PATHDSTGCP.
  - SITEENV: Decide se a copia será para o storage no dccm ou um bucket na gcp,exemplo: SITEENV="DCCM" ou SITEENV="GCP"
  - PATHSRC: Pasta de origem dos arquivos, exemplo: PATHSRC="/opt/ef-shared/...."
  - PATHDST: Pasta de destino dos arquivos para copia no DCCM, exemplo: PATHDST="/opt/ef-shared/destino/...." 
  - PATHDSTGCP: Pasta de destino dos arquivos para copia no GCP, exemplo: PATHDST="bkt-gglobo-ingest-encoding-hdg-dev-l2v
"

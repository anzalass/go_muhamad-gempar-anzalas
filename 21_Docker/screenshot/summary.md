# Docker

Docker adalah platform perangkat lunak yang digunakan untuk mengemas, mendistribusikan, dan menjalankan aplikasi dalam lingkungan yang terisolasi yang disebut "container." Container Docker mengizinkan aplikasi dan semua dependensinya untuk dijalankan dengan konsisten di berbagai lingkungan, seperti mesin pengembangan, produksi, atau cloud, tanpa perlu mengkhawatirkan perbedaan konfigurasi atau lingkungan host. Ini membuat pengembangan dan implementasi aplikasi lebih cepat, konsisten, dan mudah dikelola.

## Perintah pada docker

- `docker ps -a` -> digunakan untuk menampilkan daftar seluruh container
- `docker images -a` -> digunakan untuk melihat daftar images yg sudah di download
- `docker start <containerID/container_name>` -> digunakan untuk menjalankan container (dari kondisi stop)
- `docker stop <containerID/container_name>` -> digunakan untuk menghentikan container (dari kondisi berjalan)
- `docker rm <containerID/container_name>`-> digunakan untuk menghapus container (dari kondisi stop)
- `docker rmi <containerID>` -> digunakan untuk menghapus images (khusus image yang sedang tidak digunakan)
- `docker build -t <nama_image> .` -> digunakan untuk membuat docker image dari project. Project wajib menyertakan **dockerfile**
- `docker run <nama_image> ` -> digunakan untuk membuat container. Jika image tidak ditemukan pada local storage. Otomatis download dari dockerhub
- `docker logs <containerID/container_name>` -> digunakan untuk menampilkan log dari container terkait.

## Langkah deploy project ke server

- Clone project dari github (di sarankan ada folder khusus sebagi lokasi download project)
- buat file '.env' (`touch .env`)
- jalankan perintah `docker build -t <nama_image> .`
- <details>
      <summary> Buat Database Instance </summary>
      <details>
        <summary> Google SQL </summary>
        <ul>
          <li>Pada cloud console pilih menu SQL</li>
          <li>Create instance sesuai dengan spesifikasi dan kebutuh yang diinginkan</li>
          <li>Setelah instance telah terbuat, pilih instance, lalu masuk ke menu connections</li>
          <li>Pilih tab `networking`</li>
          <li>Pada bagian `Authorized networks`, tambahkan public IP dari Instance Compute Engine</li>
        </ul>
      </details>
      <details>
        <summary> Amazone RDS </summary>
        <ul>
          <li>Cari fitur RDS</li>
          <li>Pada bagian Databases, pilih create instance</li>
          <li>Atur spesifikasi sesuai keinginan</li>
          <li>Pada tab `security group` pilih opsi untuk mengkaitkan dengan intance EC2</li>
        </ul>
      </details>
    </details>
- Jalankan `docker run -it mysql mysql -h<publicIP/endpoint> -u<root/sesuaiPendaftarAWS> -p` perintah ini digunakan untuk login ke DB instance yg sudah dibuat, untuk create DB baru. Setelah dijalankan, perintah ini akan meniggalkan sebuah container mysql dalam kondisi stop. Container tsb bisa digunakan kembali dengan cara `docker exec -it <containerID/container_name> bash`, lakukan login mysql dengan perintah CLI. untuk keluar jalankan perintah `exit`
- Jalankan perintah `docker run 
--name <container_name> 
-p 8000:8000 
-d 
-e DBHOST=<publicIP/endpoint> 
-e DBPORT=3306 
-e DBUSER=<root/sesuaiPendaftarAWS> 
-e DBNAME=<sesuaikanDengan pembuatan> 
-e DBPASS=<sesuaikanWaktuBikinDB> 
-e SECRET=S3creT 
-e REFSECRET=rEFSCret 
-e SERVER=8000 <nama_image_project>`

# Compute Service

Compute Service adalah layanan komputasi yang disediakan dalam komputasi awan (cloud computing). Layanan ini memungkinkan pengguna untuk menyewa dan menggunakan sumber daya komputasi seperti server virtual, CPU, RAM, dan penyimpanan yang di-host di infrastruktur cloud penyedia layanan. Dengan menggunakan Compute Service, pengguna dapat membuat, mengelola, dan mengubah server virtual, menginstal dan menjalankan perangkat lunak, serta menangani beban kerja komputasi mereka.

Contoh dari Compute Service yang populer adalah Amazon Elastic Compute Cloud (EC2) dari AWS, Google Compute Engine dari Google Cloud, dan Azure Virtual Machines dari Microsoft Azure. Layanan-layanan ini memungkinkan pengguna untuk dengan mudah menyediakan dan mengelola sumber daya komputasi sesuai kebutuhan mereka, baik untuk keperluan pengembangan, hosting aplikasi web, analisis data, atau keperluan lainnya dalam lingkungan awan.

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

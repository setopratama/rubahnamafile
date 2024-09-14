# Aplikasi Pengubah Nama File

Aplikasi ini adalah program sederhana yang ditulis dalam bahasa Go untuk mengubah nama file dalam suatu folder dengan menambahkan prefix dan nomor urut.

## Cara Menjalankan Proyek

1. Pastikan Anda telah menginstal Go di komputer Anda. Jika belum, Anda dapat mengunduhnya dari [situs resmi Go](https://golang.org/dl/).

2. Unduh atau clone repositori ini ke komputer Anda.

3. Buka terminal atau command prompt, lalu navigasikan ke direktori proyek.

4. Jalankan program dengan perintah berikut:

   ```
   go run main.go
   ```

5. Program akan meminta Anda untuk memasukkan:
   - Prefix untuk nama file (contoh: FT)
   - Path folder yang ingin diubah nama filenya (contoh: C:\Users\NamaAnda\Documents\FolderUntukDiubah atau /home/namaanda/documents/folderuntukdiubah)

6. Setelah Anda memasukkan informasi yang diminta, program akan mengubah nama file dalam folder yang ditentukan dengan menambahkan prefix dan nomor urut.

7. Program akan menampilkan hasil perubahan nama file di layar.

## Catatan Penting

- Pastikan Anda memiliki izin untuk mengubah file di folder yang ditentukan.
- Program ini akan mengabaikan subfolder dan hanya mengubah nama file di tingkat atas folder yang ditentukan.
- File akan diurutkan berdasarkan nama sebelum diubah namanya.
- Ekstensi file akan tetap dipertahankan setelah perubahan nama.

Selamat mencoba!

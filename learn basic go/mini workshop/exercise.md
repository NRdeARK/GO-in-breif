
## แบบฝึกหัด

## แบบฝึกหัด: Println

โจทย์ปัญหา: https://go.dev/play/p/WPtOfb_YsS3

```
เรื่อง: Avengers: Endgame
ปี: 2019
เรตติ้ง: 8.4
ประเภท: Sci-Fi
ซุปเปอร์ฮีโร่: true
เรื่องโปรด: ⭐
```

- เฉลย 1 literal: https://go.dev/play/p/jhfzQ9jjsin
- เฉลย 2 var: https://go.dev/play/p/PQHt1WbwyL0
- เฉลย 3 raw string `` : https://go.dev/play/p/E6prOIF414E

## แบบฝึกหัด: Printf

โจทย์ปัญหา: https://go.dev/play/p/0WRuhh-jFY1

```
เรื่อง: Avengers: Endgame
ปี: 2019
เรตติ้ง: 8.4
ประเภท: Sci-Fi
ซุปเปอร์ฮีโร่: true
เรื่องโปรด: ⭐
```

- เฉลย 1 joke raw string ``: https://go.dev/play/p/Vx6DEyRL_Wx
- เฉลย 2 joke: https://go.dev/play/p/j40F8V2pSn9

## แบบฝึกหัด: if

โจทย์ปัญหา: https://go.dev/play/p/5qBXcXFYOJ5

- 1. ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการแสดงผลคำว่า "Disappointed 😞"
- 2. ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการแสดงผลคำว่า "Normal 😐"
- 3. ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการแสดงผลคำว่า "Good 🥰"
- 4. กรณีอื่นๆ ให้ทำการแสดงผลคำว่า "🤷🤷🤷🤷"

* เฉลย 1: https://go.dev/play/p/ZsnHXfHZmNh

## แบบฝึกหัด: Switch case

โจทย์ปัญหา: https://go.dev/play/p/AKlYvUycUWj

- 1. ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการแสดงผลคำว่า "Disappointed 😞"
- 2. ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการแสดงผลคำว่า "Normal 😐"
- 3. ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการแสดงผลคำว่า "Good 🥰"
- 4. กรณีอื่นๆ ให้ทำการแสดงผลคำว่า "🤷🤷🤷🤷"

* เฉลย 1: https://go.dev/play/p/UmtyGq2vuKI

## แบบฝึกหัด: function

โจทย์ปัญหา: https://go.dev/play/p/Lpl1dtmHbXn

- 1. สร้างฟังก์ชันชื่อ emote และรับพารามิเตอร์หนึ่งตัวชื่อ ratings เป็นตัวแปรชนิด float64 และคืนค่าเป็น string
- 2. ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการคืนค่าคำว่า "Disappointed 😞"
- 3. ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการคืนค่าคำว่า "Normal 😐"
- 4. ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการคืนค่าคำว่า "Good 🥰"
- 5. กรณีอื่นๆ ให้ทำการคืนค่าคำว่า "🤷🤷🤷🤷"

* เฉลย 1: https://go.dev/play/p/b8QAB5cJVlc

## แบบฝึกหัด: array

โจทย์ปัญหา: https://go.dev/play/p/ynSrKmKWBt-

- 1. ให้ประกาศตัวแปรอาร์เรย์ประเภทหนัง(genres)ที่เก็บค่าเป็น "Action", "Adventure" และ "Fantasy" ตามลำดับ
- 2. ให้อ่านค่าของอาร์เรย์แต่ละช่องเพื่อแสดงผล ให้แสดงผลทั้งหมดตามตัวอย่าง Output ด้านล่าง
- 3. หลังจากนั้นเปลี่ยนแปลงค่าในอาร์เรย์ index ที่ 1 ให้เป็น Sci-Fi พร้อมทั้งแสดงผล เพื่อยืนยันว่าค่าเปลี่ยนจริง

Output:
"genres[0]: Action"
"genres[1]: Adventure"
"genres[2]: Fantasy"
"genres[1]: Sci-Fi"

- เฉลย 1: https://go.dev/play/p/MC8ICOf4rqy
- เฉลย 2 [...]: https://go.dev/play/p/Xk2PmBw6nVk

## แบบฝึกหัด: for, for-range

โจทย์ปัญหา: https://go.dev/play/p/MCDLMUHsoVV

- 1. ให้ใช้ for loop ทำการเปลี่ยนค่าในอาร์เรย์ genres โดยเพิ่มคำนำหน้า "Movie: " ทุกค่า ดังนี้ "Movie: Action", "Movie: Adventure", "Movie: Fantasy"
- 2. ให้แสดงผลค่า genres ทีละค่า โดยใช้ for-range

Output:
before for loop: [3]string{"Action", "Adventure", "Fantasy"}
after for loop: [3]string{"Movie: Action", "Movie: Adventure", "Movie: Fantasy"}
genres[0]: Movie: Action
genres[1]: Movie: Adventure
genres[2]: Movie: Fantasy

- เฉลย 1: https://go.dev/play/p/dXS8IwBEmMe

## แบบฝึกหัด: slice

โจทย์ปัญหา: https://go.dev/play/p/C90FxMb1bQm

- 1. เราได้เก็บสะสมคะแนนโหวตผู้ชมไว้เป็น 2 ชุด ที่เก็บอยู่ในตัวแปร xs และ ys เรียบร้อยแล้ว
- 2. ให้ทำการรวมคะแนนโหวตที่อยู่ในตัวแปร xs และ ys ไปเก็บไว้ในตัวแปร votes ตามลำดับ (ค่าใน xs ทั้งหมดแล้วต่อด้วย ys).
- 3. ทำการแสดงผลคะแนนโหวตของผู้ชมที่อยู่ในตำแหน่ง(index)ที่ 5 ถึง 8 ของ votes ออกมาทางหน้าจอ

* เฉลย: https://go.dev/play/p/STHSciSQMb4

## แบบฝึกหัด: struct

โจทย์ปัญหา: https://go.dev/play/p/GsAErtuEKDD

- 1. ให้นิยามโครงสร้างข้อมูล movie เพื่อเก็บ ชื่อเรื่อง(string) ปี(ตัวเลข) เรตติ้ง(ตัวเลขทศนิยม) ประเภท(slice ของ string) และ isSuperHero(bool)
- 2. ให้ประกาศตัวแปรเพื่อเก็บหนัง Avengers: Endgame, ปี:2019, เรตติ้ง:8.4, ประเภท:["Action", "Drama"] และ isSuperHero:true
- 3. ให้ประกาศตัวแปรเพื่อเก็บหนัง Infinity War, ปี:2018, เรตติ้ง:8.4, ประเภท:["Action", "Sci-Fi"] และ isSuperHero:true
- 4. ให้เก็บหนังทั้งสองเรื่องไว้ในตัวแปร mvs 5. ทำการวนลูปเพื่อแสดงผลหนังทีละเรื่อง

Output:
main.movie{title:"Avengers: Endgame", year:2019, rating:8.4, genres:[]string{"Action", "Drama"}, isSuperHero:true}
main.movie{title:"Avengers: Infinity War", year:2018, rating:8.4, genres:[]string{"Action", "Sci-Fi"}, isSuperHero:true}

- เฉลย: https://go.dev/play/p/2yEHW0R92M8

## แบบฝึกหัด: method

โจทย์ปัญหา: https://go.dev/play/p/cOEbj0wB9DO

- เฉลย 1: https://go.dev/play/p/gErFHYR3dhY

## แบบฝึกหัด: pointer of struct

โจทย์ปัญหา: https://go.dev/play/p/dt3Y7_4hqSy

- 1. ให้สร้าง method addVote รับพารามิเตอร์ rating เป็น float64
- 2. ให้ method addVote เพิ่มค่า rating เข้าไปใน slice votes ของ struct movie

* เฉลย 1: https://go.dev/play/p/57CLX56yQ__f

## แบบฝึกหัด: interface

โจทย์ปัญหา: https://go.dev/play/p/Edrx94i72PK

- เฉลย 1: https://go.dev/play/p/53uNT-ReSgG

## แบบฝึกหัด: maps

โจทย์ปัญหา: https://go.dev/play/p/Fgvv6ad76DJ

- เฉลย 1: https://go.dev/play/p/qK2wa358ILD

## แบบฝึกหัด: struct json

โจทย์ปัญหา: https://go.dev/play/p/tbEr1wewOwL

- เฉลย 1: https://go.dev/play/p/tavKPCcpnDG

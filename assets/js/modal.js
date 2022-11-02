const buttonOpenLogin = document.getElementById('modalOpenLogin');
const buttonOpen01 = document.getElementById('modalOpen01');
const buttonOpen02 = document.getElementById('modalOpen02');
const modal01 = document.getElementById('signUpModal');
const modal02 = document.getElementById('loginModal');
const buttonClose = document.getElementsByClassName('modalCloseButton');

// 会員登録ボタンがクリックされた時
buttonOpen01.addEventListener('click', modalOpen);
buttonOpen02.addEventListener('click', modalOpen);
function modalOpen() {
  modal01.style.display = 'block';
}

// ログインボタンがクリックされた時
buttonOpenLogin.addEventListener('click', modalOpenLogin);
function modalOpenLogin() {
 modal02.style.display = 'block';
}

// バツ印がクリックされた時
for(var i = 0; i < buttonClose.length; i++) {
    buttonClose[i].addEventListener('click', modalClose);
}
function modalClose() {
  modal01.style.display = 'none';
  modal02.style.display = 'none';
}

// モーダルコンテンツ以外がクリックされた時
addEventListener('click', outsideClose);
function outsideClose(e) {
  if (e.target == modal01) {
    modal01.style.display = 'none';
  }
  else if (e.target == modal02) {
    modal02.style.display = 'none';
  }
}
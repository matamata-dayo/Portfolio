// documentを取得
let selectTeam = document.getElementById('select-box-team');
let selectPosition = document.getElementById('select-box-position');
let selectAge = document.getElementById('select-box-age');
let selectCountry = document.getElementById('select-box-country');
let selectNumber = document.getElementById('select-box-number');
let selectCaption = document.getElementById('select-box-captain');
let selectFoot = document.getElementById('select-box-foot');
let selectHeight = document.getElementById('select-box-height');
let buttonReset = document.getElementById('inputReset');

// セレクトボックスが選択された際、画像を変更
selectTeam.onchange = function() {
  document.getElementById('team-img').src = '../../assets/img/selectBox/teams/' + selectTeam.value + '.png';
}

selectPosition.onchange = function() {
  document.getElementById('position-img').src = '../../assets/img/selectBox/positions/' + selectPosition.value + '.png';
}

selectCountry.onchange = function() {
  document.getElementById('country-img').src = '../../assets/img/selectBox/countries/' + selectCountry.value + '.png';
  document.getElementById('country-img').style.border = '2px solid #000000';
}

selectAge.onchange = function() {
  document.getElementById('age-img').src = '../../assets/img/selectBox/age.png';
}

selectNumber.onchange = function() {
  document.getElementById('number-img').src = '../../assets/img/selectBox/number.png';
  document.getElementById('numberValue').innerHTML = selectNumber.value;
}

selectCaption.onchange = function() {
  document.getElementById('captain-img').src = '../../assets/img/selectBox/captain-band.png';
}

selectFoot.onchange = function() {
  document.getElementById('foot-img').src = '../../assets/img/selectBox/foot/' + selectFoot.value + '.png';
}

selectHeight.onchange = function() {
  document.getElementById('height-img').src = '../../assets/img/selectBox/height.png';
}

// リセットボタン押下時、画像を全て初期化
buttonReset.onclick = function() {
  const initialImg = '../../assets/img/selectBox/NoSelect.png';
  document.getElementById('team-img').src = initialImg;
  document.getElementById('position-img').src = initialImg;
  document.getElementById('country-img').src = initialImg;
  document.getElementById('country-img').style.border = '0px'
  document.getElementById('number-img').src = initialImg;
  document.getElementById('numberValue'). innerHTML = '';
  document.getElementById('age-img').src = initialImg;
  document.getElementById('captain-img').src = initialImg;
  document.getElementById('foot-img').src = initialImg;
  document.getElementById('height-img').src = initialImg;
}
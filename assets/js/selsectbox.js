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

// 変更後の画像
const initialImg = '../../assets/img/selectBox/NoSelect.png'

const teamImg = [
  '../../assets/img/selectBox/teams/Arsenal.png', 
  '../../assets/img/selectBox/teams/Chelsea.png',
  '../../assets/img/selectBox/teams/Liverpool.png',
  '../../assets/img/selectBox/teams/Manchester-City.png',
  '../../assets/img/selectBox/teams/Manchester-United.png',
  '../../assets/img/selectBox/teams/Tottenham.png',
];

const positionImg = [
  '../../assets/img/selectBox/positions/FW_img.png', 
  '../../assets/img/selectBox/positions/MF_img.png',
  '../../assets/img/selectBox/positions/DF_img.png',
  '../../assets/img/selectBox/positions/GK_img.png'
];

const numberImg = '../../assets/img/selectBox/tshirt.png'

const captainImg = '../../assets/img/selectBox/captain-band.png'

const footImg = [
  '../../assets/img/selectBox/foot/right-foot.png', 
  '../../assets/img/selectBox/foot/left-foot.png'
]

// セレクトボックスが選択された際、画像を変更
selectTeam.onchange = function() {
  if (selectTeam.value === 'Arsenal') {
    document.getElementById('team-img').src = teamImg[0];
  } else if (selectTeam.value === 'Chelsea') {
    document.getElementById('team-img').src = teamImg[1];
  } else if (selectTeam.value === 'Liverpool') {
    document.getElementById('team-img').src = teamImg[2];
  } else if (selectTeam.value === 'Manchester-City') {
    document.getElementById('team-img').src = teamImg[3];
  } else if (selectTeam.value === 'Manchester-United') {
    document.getElementById('team-img').src = teamImg[4];
  } else if (selectTeam.value === 'Tottenham-Hotspur') {
    document.getElementById('team-img').src = teamImg[5];
  }
}

selectPosition.onchange = function() {
  if (selectPosition.value === 'fw') {
    document.getElementById('position-img').src = positionImg[0];
  } else if (selectPosition.value === 'mf') {
    document.getElementById('position-img').src = positionImg[1];
  } else if (selectPosition.value === 'df') {
    document.getElementById('position-img').src = positionImg[2];
  } else if (selectPosition.value === 'gk') {
    document.getElementById('position-img').src = positionImg[3];
  }
}

selectNumber.onchange = function() {
  document.getElementById('number-img').src = numberImg;
  document.getElementById('numberValue').innerHTML = selectNumber.value;
}

selectCaption.onchange = function() {
  document.getElementById('captain-img').src = captainImg;
}

selectFoot.onchange = function() {
  if (selectFoot.value === 'right') {
    document.getElementById('foot-img').src = footImg[0];
  } else {
    document.getElementById('foot-img').src = footImg[1];
  }
}

// リセットボタン押下時、画像を全て初期化
buttonReset.onclick = function() {
  document.getElementById('team-img').src =initialImg
  document.getElementById('position-img').src =initialImg
  document.getElementById('country-img').src =initialImg
  document.getElementById('number-img').src =initialImg
  document.getElementById('numberValue').innerHTML = "";
  document.getElementById('age-img').src =initialImg
  document.getElementById('captain-img').src =initialImg
  document.getElementById('foot-img').src =initialImg
  document.getElementById('height-img').src =initialImg
}
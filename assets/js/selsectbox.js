let selectTeam = document.getElementById('select-box-team');

selectTeam.onchange = function() {

  const img = [
    '../../assets/img/teams/Arsenal.jpeg', 
    '../../assets/img/teams/Chelsea.jpeg',
    '../../assets/img/teams/Liverpool.jpeg',
    '../../assets/img/teams/Manchester-city.jpeg',
    '../../assets/img/teams/Manchester-united.jpg',
    '../../assets/img/teams/Tottenham.jpeg',
    '../../assets/img/teams/FW_img.png'
  ];

  if (selectTeam.value === 'Arsenal') {
    document.getElementById('team-img').src = img[0];
  } else if (selectTeam.value === 'Chelsea') {
    document.getElementById('team-img').src = img[1];
  } else if (selectTeam.value === 'Liverpool') {
    document.getElementById('team-img').src = img[2];
  } else if (selectTeam.value === 'Manchester-City') {
    document.getElementById('team-img').src = img[3];
  } else if (selectTeam.value === 'Manchester-United') {
    document.getElementById('team-img').src = img[4];
  } else if (selectTeam.value === 'Tottenham-Hotspur') {
    document.getElementById('team-img').src = img[5];
  }
}

let selectPosition = document.getElementById('select-box-position');

selectPosition.onchange = function() {

  const img = [
    '../../assets/img/positions/FW_img.png', 
    '../../assets/img/positions/MF_img.png',
    '../../assets/img/positions/DF_img.png',
    '../../assets/img/positions/GK_img.png'
  ];

  if (selectPosition.value === 'fw') {
    document.getElementById('position-img').src = img[0];
  } else if (selectPosition.value === 'mf') {
    document.getElementById('position-img').src = img[1];
  } else if (selectPosition.value === 'df') {
    document.getElementById('position-img').src = img[2];
  } else if (selectPosition.value === 'gk') {
    document.getElementById('position-img').src = img[3];
  }
}
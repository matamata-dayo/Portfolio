let selectTeam = document.getElementById('select-box-team');

selectTeam.onchange = function() {

  const img = [
    "../../assets/img/Arsenal.jpeg", 
    "../../assets/img/Chelsea.jpeg",
    "../../assets/img/Liverpool.jpeg",
    "../../assets/img/Manchester-city.jpeg",
    "../../assets/img/Manchester-united.jpg",
    "../../assets/img/Tottenham.jpeg"
  ];

  if (selectTeam.value === "Arsenal") {
    document.getElementById('team-img').src = img[0];
  } else if (selectTeam.value === "Chelsea") {
    document.getElementById('team-img').src = img[1];
  } else if (selectTeam.value === "Liverpool") {
    document.getElementById('team-img').src = img[2];
  } else if (selectTeam.value === "Manchester-City") {
    document.getElementById('team-img').src = img[3];
  } else if (selectTeam.value === "Manchester-United") {
    document.getElementById('team-img').src = img[4];
  } else if (selectTeam.value === "Tottenham-Hotspur") {
    document.getElementById('team-img').src = img[5];
  }
}
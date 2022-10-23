$(function(){
    $('.select-box-team').change(function () {
      //選択したoptionのvalueを取得
      var val = $(this).val();
      //先頭に#を付けてvalueの値をclassに変換
      var selectTeamClass = '.' + val;
      //一度すべてのブロックを非表示にする
      $('ul li').hide();
      //選択したブロックのみを表示
      $(selectTeamClass).show();
    });
  });
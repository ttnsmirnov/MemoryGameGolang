window.addEventListener('load', function() {
    setTimeout(function() {
      var cards = document.querySelectorAll('.card');
      for (var i = 0; i < cards.length; i++) {      
        cards[i].querySelector('.img').classList.toggle('hide');
      }
    }, 2000);
  });

  var images = document.querySelectorAll('.img');
  for (var i = 0; i < images.length; i++) {
      images[i].addEventListener('click', function(event) {
         // console.log(this.className);
          if(this.className == "img hide"){
              this.className ="img show img-1";  
          }
         // console.log("className: " + this.className);
  
          // Извлекаем родительскую кнопку изображения
          var parentButton = this.closest('.card');
          // Извлекаем URL фона кнопки из переменной buttonUrl
          var buttonUrl = parentButton.style.backgroundImage.replace(/^url\(["']?/, '').replace(/["']?\)$/, '');
          // Находим последний индекс символа '/' в строке
         var lastIndex = buttonUrl.lastIndexOf('/');

        // Извлекаем подстроку, начиная с символа, следующего за последним '/'
           var filename = buttonUrl.substring(lastIndex + 1); 

         // Извлекаем эллементы под ID   flip , ers
           var nrOfFlips = document.getElementById('flip');
           var nrOfErrors = document.getElementById('ers');
        
          // Отправляем AJAX-запрос на сервер Go
          var xhr = new XMLHttpRequest();
          xhr.open('POST', '/url', true);
          xhr.setRequestHeader('Content-Type', 'application/json');
          xhr.onreadystatechange = function() {
            if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
                var response = JSON.parse(xhr.responseText);
                setTimeout(function() {
                if (response.match) {
                        //console.log("response match :" + response.match);   
                        const elementsAll = document.querySelectorAll('.img-1');
                        elementsAll.forEach((el) => {
                            el.classList.toggle('img-1');                                                                                            
                    });                                       
                }else {
                        const elements = document.querySelectorAll('.img-1');
                        //console.log("response no match :"+ response.match);   
                        elements.forEach((element) => {
                        element.classList.toggle('img-1');
                        element.classList.toggle('show');
                        element.classList.add('hide');
                        
                    });
                    
                } 
                nrOfFlips.innerHTML = response.flips;
                nrOfErrors.innerHTML = response.errs;
                //console.log("nrOfErrors :"+ response.errs); 
            }, 500); }
        };
          xhr.send(JSON.stringify({buttonUrl: filename}));
      });
  }
  
  
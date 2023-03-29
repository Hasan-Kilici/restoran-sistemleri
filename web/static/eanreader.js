function controlL(text) {
    let newtext = text;
    newtext = newtext.replace(/\$/g, "L");
    newtext = newtext.replace(/\@/g, "L");
    newtext = newtext.replace(/\-/g, "0");
    newtext = newtext.replace(/\^/g, "0");
    
    return newtext;
  }
  
  
  var _scannerIsRunning = false;
  
  function startScanner() {
    Quagga.init({
      inputStream: {
        name: "Live",
        type: "LiveStream",
        target: document.querySelector('#scanner-container'),
        constraints: {
          width: 340,
          height: 340,
          facingMode: "environment"
        },
      },
      decoder: {
        readers: [
          "code_128_reader",
          "ean_reader",
        ],
        debug: {
          showCanvas: true,
          showPatches: true,
          showFoundPatches: true,
          showSkeleton: true,
          showLabels: true,
          showPatchLabels: true,
          showRemainingPatchLabels: true,
          boxFromPatches: {
            showTransformed: true,
            showTransformedBox: true,
            showBB: true
          }
        }
      },
  
    }, function(err) {
      if (err) {
        console.log(err);
        return
      }
  
      console.log("Initialization finished. Ready to start");
      Quagga.start();
  
      _scannerIsRunning = true;
    });
  
    Quagga.onProcessed(function(result) {
      setTimeout(()=>{
      var drawingCtx = Quagga.canvas.ctx.overlay,
        drawingCanvas = Quagga.canvas.dom.overlay;
  
      if (result) {
        if (result.boxes) {
          drawingCtx.clearRect(0, 0, parseInt(drawingCanvas.getAttribute("width")), parseInt(drawingCanvas.getAttribute("height")));
          result.boxes.filter(function(box) {
            return box !== result.box;
          }).forEach(function(box) {
            Quagga.ImageDebug.drawPath(box, { x: 0, y: 1 }, drawingCtx, { color: "lime", lineWidth: 2 });
          });
        }
  
        if (result.box) {
          Quagga.ImageDebug.drawPath(result.box, { x: 0, y: 1 }, drawingCtx, { color: "#00F", lineWidth: 2 });
        }
  
        if (result.codeResult && result.codeResult.code) {
          Quagga.ImageDebug.drawPath(result.line, { x: 'x', y: 'y' }, drawingCtx, { color: 'red', lineWidth: 3 });
        }
      }
      },1000);
    });
  
  
    Quagga.onDetected(function(result) {
      setTimeout(() => {
        document.getElementById("result").innerHTML = controlL(result.codeResult.code);
        var context = new AudioContext();
        var oscillator = context.createOscillator();
        oscillator.type = "sine";
        oscillator.frequency.value = 800;
        oscillator.connect(context.destination);
        oscillator.start();
  
        setTimeout(() => {
          oscillator.stop();
        }, 100);
      });
    }, 1000);
  }
  
  
  // Start/stop scanner
  document.getElementById("btn").addEventListener("click", function() {
    if (_scannerIsRunning) {
      Quagga.stop();
    } else {
      startScanner();
    }
  }, false);
  
  function start() {
    startScanner();
  }
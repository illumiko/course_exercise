class DumpKit {
  constructor() {
    //selectors
    this.playBtn = document.querySelector(".play");
    this.pads = document.querySelectorAll(".pad");
    //beats options selector
    this.hihatBeatOptions = document.querySelector("#hihatSoundSelect");
    this.clapBeatOptions = document.querySelector("#clapSoundSelect");
    this.crashBeatOptions = document.querySelector("#crashSoundSelect");
    this.snareBeatOptions = document.querySelector("#snareSoundSelect");
    this.openhatBeatOptions = document.querySelector("#openhatSoundSelect");
    this.shakerBeatOptions = document.querySelector("#shakerSoundSelect");
    this.kickBeatOptions = document.querySelector("#kickSoundSelect");

    //beat options array
    this.beatOptions = [
      this.hihatBeatOptions,
      this.clapBeatOptions,
      this.crashBeatOptions,
      this.snareBeatOptions,
      this.openhatBeatOptions,
      this.shakerBeatOptions,
      this.kickBeatOptions,
    ];

    //final value of beats
    this.hihatBeatSelected = this.hihatBeatOptions.value;
    this.clapBeatSelected = this.clapBeatOptions.value;
    this.crashBeatSelected = this.crashBeatOptions.value;
    this.snareBeatSelected = this.snareBeatOptions.value;
    this.openhatBeatSelected = this.openhatBeatOptions.value;
    this.shakerBeatSelected = this.shakerBeatOptions.value;
    this.kickBeatSelected = this.kickBeatOptions.value;
    //tempo of the play
    this.tempo = document.querySelector(".tempo");
    this.tempoVal = this.tempo.valueAsNumber;
    //status of the beat
    this.status = "off";
    //hold the interval for beat
    this.BeatPlayerHolder;
  }

  getBeat(e) {
    //getting value of beats from options
    const hihatBeat = this.hihatBeatOptions.value,
      clapBeat = this.clapBeatOptions.value,
      crashBeat = this.crashBeatOptions.value,
      snareBeat = this.snareBeatOptions.value,
      openhatBeat = this.openhatBeatOptions.value,
      shakerBeat = this.shakerBeatOptions.value,
      kickBeat = this.kickBeatOptions.value;

    //object of values of selected option
    const beats = {
      hihat: hihatBeat,
      clap: clapBeat,
      crash: crashBeat,
      snare: snareBeat,
      openhat: openhatBeat,
      shaker: shakerBeat,
      kick: kickBeat,
    };

    //setting checking which beats option was changed and declaring the changed value
    switch (e.target.id) {
      case "hihatSoundSelect":
        this.hihatBeatSelected = `${beats.hihat}`;
        console.log(this.hihatBeatOptions);
        break;
      case "clapSoundSelect":
        this.clapBeatSelected = `${beats.clap}`;
        console.log(this.clapBeatOptions);
        break;
      case "crashSoundSelect":
        this.crashBeatSelected = `${beats.crash}`;
        break;
      case "snareSoundSelect":
        this.snareBeatSelected = `${beats.snare}`;
        console.log(this.snareBeatOptions);
        break;
      case "openhatSoundSelect":
        this.openhatBeatSelected = `${beats.openhat}`;
        break;
      case "shakerSoundSelect":
        this.shakerBeatSelected = `${beats.shaker}`;
        break;
      case "kickSoundSelect":
        this.kickBeatSelected = `${beats.kick}`;
        break;
    }
  }

  valueOfSlider(e) {
    //if (this.status == "on") {
    clearInterval(this.BeatPlayerHolder);
    this.tempoVal = parseInt(e.target.valueAsNumber);
    this.status = "off";
    this.play();
    //}
    //else {
    //clearInterval(this.BeatPlayerHolder);
    //this.tempoVal = parseInt(e.target.valueAsNumber);
    //this.play();
    //this.status = "off";
    //}
  }

  play() {
    let count = 0;
    let interval = Math.round((60 / this.tempoVal) * 1500);
    console.log(interval);
    const playFunc = () => {
      count++;
      if (count !== 9) {
        let pads = document.querySelectorAll(`.pads${count}`);
        //pads.style.transform = "scale(1.1)"
        pads.forEach((i) => {
          //visual feed back for activation of pad
          i.style.animation = `padsActive 400ms ease-in-out`;

          //visual feed back for pad finishing beat
          i.addEventListener("anmationend", () => {
            i.style.animation = "none";
          });

          //cheacking if pads are active
          if (i.classList.contains("padsActive")) {
            if (i.classList.contains("hihatSoundPad")) {
              let hihatAudio = new Audio(`${this.hihatBeatSelected}`);
              hihatAudio.play();
            } else if (i.classList.contains("clapSoundPad")) {
              let clapAudio = new Audio(`${this.clapBeatSelected}`);
              clapAudio.play();
            } else if (i.classList.contains("crashSoundPad")) {
              let crashAudio = new Audio(`${this.crashBeatSelected}`);
              crashAudio.play();
            } else if (i.classList.contains("snareSoundPad")) {
              let snareAudio = new Audio(`${this.snareBeatSelected}`);
              snareAudio.play();
            } else if (i.classList.contains("openhatSoundPad")) {
              let openhatAudio = new Audio(`${this.snareBeatSelected}`);
              openhatAudio.play();
            } else if (i.classList.contains("shakerSoundPad")) {
              let shakerAudio = new Audio(`${this.shakerBeatSelected}`);
              shakerAudio.play();
            } else if (i.classList.contains("kickSoundPad")) {
              let kickAudio = new Audio(`${this.kickBeatSelected}`);
              kickAudio.play();
            }
          }
        });

        //console.log(pads);
      } else {
        count = 0;
        //let pads = document.querySelectorAll(`.pads${count}`);
        //console.log(pads);
      }
    };

    if (this.status == "on") {
      console.log("x");
      clearInterval(this.BeatPlayerHolder);
      this.status = "off";
    } else {
      this.BeatPlayerHolder = setInterval(playFunc, interval);
      this.status = "on";
    }
    /*
    if (this.status == 'off'){
      this.status = 'on'
      let playInterval = setInterval(playFunc, 1000)
      //console.log('hello')
    }
    else if (this.status == 'on') {
      this.status = 'off'
      clearInterval(playInterval)
      //console.log('bye')
    }
    */
  }
}

const drumKit = new DumpKit();

drumKit.beatOptions.forEach((i) => {
  //console.log(i);
  i.addEventListener("change", (e) => {
    drumKit.getBeat(e);
  });
});

drumKit.pads.forEach((i) => {
  i.addEventListener("click", (e) => {
    if (e.target.classList.contains("padsActive")) {
      e.target.classList.remove("padsActive");
    } else {
      e.target.classList.add("padsActive");
    }
  });
});

drumKit.playBtn.addEventListener("click", () => {
  //drumKit.play()
  drumKit.play();
});

drumKit.tempo.addEventListener("input", (e) => {
  document.querySelector(`.sliderText`).innerHTML = e.target.value;
});

drumKit.tempo.addEventListener("change", (e) => {
  drumKit.valueOfSlider(e);
});

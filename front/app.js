// タイマーのデータと機能をまとめたオブジェクト

const startbtn = document.querySelector("#start-btn");
console.log("ボタンの確認：", startbtn);
startbtn.addEventListener("click", () => {
  studyTimer.start();
});

const stopbtn = document.querySelector("#stop-btn");

stopbtn.addEventListener("click", () => {
  studyTimer.stop();
});

const resetbtn = document.querySelector("#reset-btn");

resetbtn.addEventListener("click", () => {
  studyTimer.reset();
});

const audio = new Audio("アナウンス前のピンポンパンポン.mp3");

const studylog = document.querySelector(`#log-list`);
const studyTimer = {
  timeLeft: 1500, // 初期の残り時間
  timerInterval: null, // タイマーのID

  // 1. タイマーを開始するメソッド
  start() {
    const input_data = document.querySelector("#subject-input");
    const display_time = document.querySelector("#time-display");

    if (input_data.value === "") {
      alert("科目を入れてください");
      return -1;
    }

    if (this.timerInterval) {
      return -1;
    }
    this.timerInterval = setInterval(() => {
      this.timeLeft -= 1;

      const minute = this.timeLeft / 60;

      const second = this.timeLeft % 60;

      display_time.textContent = `${Math.floor(minute)}:${second.toString().padStart(2, "0")}`;

      console.log("残り時間:", this.timeLeft);

      if (this.timeLeft === 0) {
        clearInterval(this.timerInterval);
        audio.play();
        this.savelog(input_data);
        this.rest();
      }
    }, 10);
  },

  // タイマーを止めるメソッド
  stop() {
    clearInterval(this.timerInterval);
    console.log("タイマー停止。残り：", this.timeLeft);
    this.timerInterval = null;
  },
  //タイマーをリセットするメソッド
  reset() {
    clearInterval(this.timerInterval);
    const display_time = document.querySelector("#time-display");
    display_time.textContent = "25:00";
    console.log("リセットします. タイムの初期化", (this.timeLeft = 1500));
    this.timerInterval = null;
  },
  //休憩メソッド(5分間)
  rest() {
    clearInterval(this.timerInterval);
    this.timerInterval = null;

    this.timeLeft = 300;

    const display_time = document.querySelector("#time-display");

    display_time.textContent = "5:00";

    this.timerInterval = setInterval(() => {
      console.log(this.timeLeft);

      const minute = this.timeLeft / 60;

      const second = this.timeLeft % 60;

      display_time.textContent = `${Math.floor(minute)}:${second.toString().padStart(2, "0")}`;

      this.timeLeft -= 1;

      if (this.timeLeft === 0) {
        audio.play();
        display_time.textContent = "お疲れさまでした！";
        clearInterval(this.timerInterval);
      }
    }, 10);
  },

  async savelog(input_data) {
    const dataToSend = {
      subject: input_data.value,
      study_time: 25,
    };
    studylog.append(input_data.value);

    try {
      const response = await fetch("/api/save", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(dataToSend),
      });

      if (response.ok) {
        console.log("成功");
      } else {
        console.log("サーバーエラー", response.status);
        alert("保存できませんでした");
      }
    } catch (error) {
      console.log("接続失敗", error);
    }
  },
};

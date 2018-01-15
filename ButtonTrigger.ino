const int buttonPin = 2;

int buttonState = 0;
int oldState = 0;

void setup() {
  pinMode(buttonPin, INPUT);
  Serial.begin(9600);
}

void loop() {
  buttonState = digitalRead(buttonPin);
  if (oldState != buttonState) {
    if (buttonState == LOW) {
      Serial.print("RELEASE THE KRAKEN!\n");
    }
    oldState = buttonState;
  }
  delay(100);
}

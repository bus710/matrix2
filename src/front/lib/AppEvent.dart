abstract class AppEvent {
  bool type;
}

class SwitchEvent extends AppEvent {
  bool type;

  SwitchEvent(type){
    this.type = type;
  }
}


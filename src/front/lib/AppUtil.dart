import 'package:flutter_web/material.dart';

class MyPainter extends CustomPainter {
  bool state = false;

  MyPainter(bool state) {
    // this.state = false;
    this.state = state;
  }

  @override
  void paint(Canvas canvas, Size size) {
    Color selectedColor = Colors.white;
    Offset start1 = new Offset(14, 15);
    Offset end1 = new Offset(29, 30);
    Offset start2 = new Offset(14, 30);
    Offset end2 = new Offset(29, 15);

    // double radius = 10;
    Paint brush = new Paint()
      ..color = selectedColor
      ..style = PaintingStyle.stroke
      ..strokeCap = StrokeCap.round
      ..strokeWidth = 5;

    if (this.state == true) {
      canvas.drawLine(start1, end1, brush);
      canvas.drawLine(start2, end2, brush);
    }
  }

  @override
  bool shouldRepaint(CustomPainter oldDelegate) {
    return true;
  }
}

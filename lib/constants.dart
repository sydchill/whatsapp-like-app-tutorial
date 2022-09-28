import 'package:flutter/material.dart';

const textFieldStyle = TextStyle(
    fontFamily: 'Poppins',
    fontSize: 15,
    fontWeight: FontWeight.w600,
    color: Colors.black);
InputDecoration inputDecoration(label) {
  return InputDecoration(
      border: OutlineInputBorder(
          borderSide: const BorderSide(color: Colors.black),
          borderRadius: BorderRadius.circular(5)),
      labelText: label,
      labelStyle: textFieldStyle);
}

// ignore_for_file: avoid_unnecessary_containers, prefer_const_constructors

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:whatslikeapp/constants.dart';

class Authenticate extends StatefulWidget {
  const Authenticate({Key? key}) : super(key: key);

  @override
  State<Authenticate> createState() => _AuthenticateState();
}

class _AuthenticateState extends State<Authenticate> {
  bool isSignIn = true;

  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final TextEditingController _username = TextEditingController();
  final TextEditingController _password = TextEditingController();
  final TextEditingController _confirm = TextEditingController();
  final TextEditingController _email = TextEditingController();

  void _switch() {
    isSignIn = !isSignIn;
    FocusScope.of(context).unfocus();
    clearTextInput();
  }

  clearTextInput() {
    _username.clear();
    _email.clear();
    _password.clear();
    _confirm.clear();
  }

  @override
  Widget build(BuildContext context) {
    Widget page = isSignIn ? _signIn() : _signUp();

    return Scaffold(
      body: AnnotatedRegion<SystemUiOverlayStyle>(
        value: const SystemUiOverlayStyle(statusBarColor: Colors.black),
        child: GestureDetector(
          onTap: () => FocusScope.of(context).unfocus(),
          child: SingleChildScrollView(
            physics: const AlwaysScrollableScrollPhysics(),
            child: Form(
              key: _formKey,
              child: Container(
                  padding: const EdgeInsets.fromLTRB(40, 60, 40, 20),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    children: [
                      Text(
                        'Hi.Com',
                        textAlign: TextAlign.center,
                        style: TextStyle(fontSize: 45, fontFamily: 'Baloo 2'),
                      ),
                      SizedBox(height: 60),
                      page
                    ],
                  )),
            ),
          ),
        ),
      ),
    );
  }

  Widget _signIn() {
    return Container(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          _buildUsername(),
          _buildPassword(),
          SizedBox(height: 25),
          ElevatedButton(
            onPressed: () => {},
            style: ButtonStyle(
                backgroundColor:
                    MaterialStateProperty.all<Color>(Colors.black)),
            child: const Text(
              'Sign In',
              style: TextStyle(
                  fontFamily: 'Poppins',
                  fontSize: 18,
                  fontWeight: FontWeight.w400),
            ),
          ),
          SizedBox(height: 20),
          Text(
            'Create An Account',
            style: TextStyle(fontFamily: 'Poppins'),
            textAlign: TextAlign.center,
          ),
          TextButton(
            onPressed: () {
              setState(() {
                _switch();
              });
            },
            child: Text(
              'Sign Up',
              style: TextStyle(
                  color: Colors.black,
                  fontSize: 20,
                  fontWeight: FontWeight.w800,
                  fontFamily: 'Poppins'),
            ),
          )
        ],
      ),
    );
  }

  Widget _signUp() {
    return Container(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          _buildUsername(),
          _buildEmail(),
          _buildPassword(),
          _buildConfirm(),
          SizedBox(height: 25),
          ElevatedButton(
            onPressed: () => {},
            style: ButtonStyle(
                backgroundColor:
                    MaterialStateProperty.all<Color>(Colors.black)),
            child: const Text(
              'Sign Up',
              style: TextStyle(
                  fontFamily: 'Poppins',
                  fontSize: 18,
                  fontWeight: FontWeight.w400),
            ),
          ),
          SizedBox(height: 20),
          Text(
            'Already have an Account',
            style: TextStyle(fontFamily: 'Poppins'),
            textAlign: TextAlign.center,
          ),
          TextButton(
            onPressed: () {
              setState(() {
                _switch();
              });
            },
            child: Text(
              'Sign Up',
              style: TextStyle(
                  color: Colors.black,
                  fontSize: 20,
                  fontWeight: FontWeight.w800,
                  fontFamily: 'Poppins'),
            ),
          )
        ],
      ),
    );
  }

  Widget _buildUsername() {
    return Padding(
      padding: EdgeInsets.symmetric(vertical: 10),
      child: TextFormField(
        controller: _username,
        decoration: inputDecoration('Username'),
        validator: (String? value) {
          if (value!.isEmpty) {
            return 'Username is required';
          }
          return null;
        },
      ),
    );
  }

  Widget _buildPassword() {
    return Padding(
      padding: EdgeInsets.symmetric(vertical: 10),
      child: TextFormField(
        controller: _password,
        obscureText: true,
        decoration: inputDecoration('Password'),
        validator: (String? value) {
          if (value!.isEmpty) {
            return 'password is required';
          }
          return null;
        },
      ),
    );
  }

  Widget _buildConfirm() {
    return Padding(
      padding: EdgeInsets.symmetric(vertical: 10),
      child: TextFormField(
        controller: _confirm,
        decoration: inputDecoration('Confirm password'),
        obscureText: true,
        validator: (String? value) {
          if (value!.isEmpty) {
            return 'Username is required';
          }
          return null;
        },
      ),
    );
  }

  Widget _buildEmail() {
    return Padding(
      padding: EdgeInsets.symmetric(vertical: 10),
      child: TextFormField(
        controller: _email,
        decoration: inputDecoration('Email'),
        obscureText: true,
        validator: (String? value) {
          if (value!.isEmpty) {
            return 'confirm password is required';
          }
          return null;
        },
      ),
    );
  }
}

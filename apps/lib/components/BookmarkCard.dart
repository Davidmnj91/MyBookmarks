import 'package:flutter/material.dart';

class BookmarkCard extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Card(
        clipBehavior: Clip.antiAlias,
        child: Row(children: <Widget>[
          Expanded(
              flex: 1, child: Image.network('https://picsum.photos/150/150')),
          Expanded(
            flex: 2,
            child: Container(
                height: 130.0,
                padding: EdgeInsets.symmetric(horizontal: 10),
                child: Column(
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    mainAxisAlignment: MainAxisAlignment.start,
                    children: <Widget>[
                      Text("Titleee",
                          style: TextStyle(
                              fontSize: 20.0, fontWeight: FontWeight.bold)),
                      Container(height: 5),
                      Text("Loooooooooooooong subtitle",
                          style: TextStyle(fontSize: 14.0))
                    ])),
          ),
        ]));
  }
}

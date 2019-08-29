
var MYTELEOP = {
  scale: 1.0,

  Listen: function (options) {
    var that = this;
    options = options || {};
  
    var ros = options.ros;
    var topic = options.topic || '/cmd_vel';
  
    var x = 0;
    var y = 0;
    var z = 0;
  
    var cmdVel = new ROSLIB.Topic({
      ros: ros,
      name: topic,
      messageType: 'geometry_msgs/Twist'
    });
  
    var handleKey = function (key, keyDown) {
      var oldX = x;
      var oldY = y;
      var oldZ = z;
  
      var pub = true;
  
      var speed = 0;
      if (keyDown === true) {
        speed = that.scale;
      }

      switch (key) {
        case 'w':
          x = 0.5 * speed;
          break;
        case 's':
          x = -0.5 * speed;
          break;
        case 'a':
          z = 1 * speed;
          break;
        case 'd':
          z = -1 * speed;
          break;
        case 'e':
          y = -0.5 * speed;
          break;
        case 'q':
          y = 0.5 * speed;
          break;
        default:
          pub = false;
      }

      if (pub === true) {
        var twist = new ROSLIB.Message({
          angular: {
            x: 0,
            y: 0,
            z: z
          },
          linear: {
            x: x,
            y: y,
            z: z
          }
        });
        cmdVel.publish(twist);
  
        if (oldX !== x || oldY !== y || oldZ !== z) {
          that.emit('change', twist);
        }
      }
    };
  
    // handle the key
    var body = document.getElementsByTagName('body')[0];
    body.addEventListener('keydown', function (e) {
      console.log(`keydown event "${e.key}"`)
      handleKey(e.key, true);
    }, false);
    body.addEventListener('keyup', function (e) {
      console.log(`keyup event "${e.key}"`)
      handleKey(e.key, false);
    }, false);
  }
};

MYTELEOP.Listen.prototype.__proto__ = EventEmitter2.prototype;

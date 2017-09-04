//
//  NetworkController.h
//
//  Copyright (c) 2012 Peter Bakhirev <peter@byteclub.com>
//
//  Permission is hereby granted, free of charge, to any person
//  obtaining a copy of this software and associated documentation
//  files (the "Software"), to deal in the Software without
//  restriction, including without limitation the rights to use,
//  copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the
//  Software is furnished to do so, subject to the following
//  conditions:
//  
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//  
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
//  OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//  HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
//  WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
//  FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
//  OTHER DEALINGS IN THE SOFTWARE.

#import <Foundation/Foundation.h>


// Block typedefs
@class NetworkController;
typedef void (^ConnectionBlock)(NetworkController*);
typedef void (^MessageBlock)(NetworkController*,NSString*);


@interface NetworkController : NSObject<NSStreamDelegate> {
  // Connection info
  NSString* host;
  int port;
  
  // Input
  NSInputStream* inputStream;
  NSMutableData* inputBuffer;
  BOOL isInputStreamOpen;
  
  // Output
  NSOutputStream* outputStream;
  NSMutableData* outputBuffer;
  BOOL isOutputStreamOpen;
  
  // Event handlers
  MessageBlock messageReceivedBlock;
  ConnectionBlock connectionOpenedBlock;
  ConnectionBlock connectionFailedBlock;
  ConnectionBlock connectionClosedBlock;
}

// Singleton instance
+ (NetworkController*)sharedInstance;

// Methods
- (void)connect;
- (void)disconnect;
- (void)sendMessage:(NSString*)message;

// Properties
@property (copy) MessageBlock messageReceivedBlock;
@property (copy) ConnectionBlock connectionOpenedBlock;
@property (copy) ConnectionBlock connectionFailedBlock;
@property (copy) ConnectionBlock connectionClosedBlock;

@end

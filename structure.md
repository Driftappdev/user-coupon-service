🔄 Flow หลักของระบบ

1. 📥 รับ Coupon จาก Coupon Service (ผ่าน NATS)
Flow:
NATS ส่ง event มา (coupon.created / updated / redeemed)
coupon_event_consumer.go รับ event
ส่งเข้า CouponEventHandler
handler เรียก CouponCommandService
👉 service นี้ “sync coupon มาเก็บไว้ใน DB ของตัวเอง”
----------------------------------------------------------
2. 📥 รับ Order Completed (ใช้ coupon สำเร็จ)
Flow:
event: order.completed
order_completed_consumer.go รับ
แปลง DTO
ส่งเข้า OrderCompletedHandler
เรียก: RedeemCoupon(userID, couponCode)
👉 เมื่อ order สำเร็จ → coupon ถูกใช้ → update usage
----------------------------------------------------------
3. 🌐 HTTP API (Frontend ใช้)
Endpoint:
GET /api/v1/user-coupons?user_id=xxx
Flow:
Frontend → HTTP → QueryService → DB
👉 ดึง coupon ทั้งหมดของ user
----------------------------------------------------------
4. ⚡ gRPC (Order Service เรียกก่อน checkout)
Method:
Flow:
Order Service → gRPC → QueryService → DB
Logic:

      1. หา coupon
      2. เช็ค CanUse()
      3. return valid / error message
👉 ใช้ validate ก่อนกดจ่ายเงิน
----------------------------------------------------------
5. ⏰ Worker (Expire Coupon)
ทำงานทุก : 10 นาที
Flow:
Worker → CommandService → Repository
Logic:
      UPDATE user_coupons
      SET status='expired'
      WHERE valid_to < now()
----------------------------------------------------------
🔌 การเชื่อมต่อกับ Service อื่น
Service	      วิธีเชื่อม	      ใช้ทำอะไร
Coupon Service	NATS	        ส่ง coupon มาให้
Order Service	NATS	        แจ้งว่าใช้ coupon แล้ว
Order Service	gRPC	        validate coupon ก่อน checkout
Frontend	      HTTP	        แสดง coupon


🔁 Flow ใหญ่สุด (End-to-End)
🧾 แจก coupon
Coupon Service
   ↓
NATS (coupon.created)
   ↓
user-coupon-service
   ↓
AssignCoupon → DB
----------------------
🛒 ใช้ coupon
Frontend → Order Service
           ↓
     gRPC ValidateCoupon
           ↓
     OK → ทำ order
           ↓
Order Completed Event
           ↓
user-coupon-service
           ↓
----------------------
⏰ coupon หมดอายุ
Worker (ทุก 10 นาที)
   ↓
Expire()
   ↓
DB update           
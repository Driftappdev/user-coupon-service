(1) UI เปิดหน้า
→ gRPC → Query → ได้ list คูปอง

(2) user เลือกคูปอง
→ gRPC → Validate → เช็คใช้ได้ไหม

(3) user กดใช้
→ gRPC → Reserve → lock คูปองไว้

(4) Order Service ทำการจ่ายเงิน

    ├── ถ้าสำเร็จ
    │     → emit event → order_completed
    │     → Usage Service (ใช้จริง)
    │
    └── ถ้าล้มเหลว
          → emit event → order_failed
          → Release Service (คืนคูปอง)



ส่วน	            Trigger	        ทำอะไร
Query Service	    gRPC	       ดู + validate
Reserve Service	    gRPC	       จองคูปอง
Usage Service	    Event	       ใช้จริง
Release Service	    Event	       คืนคูปอง
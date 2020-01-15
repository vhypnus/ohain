package com.caocao.hermes.order.dao.model;

import com.alibaba.fastjson.annotation.JSONField;
import com.fasterxml.jackson.annotation.JsonIgnore;
import lombok.Data;

/**
 * 公务订单的账单模型，冗余计费账单的模型
 */
@Data
public class OsBillInfo extends Model {
    private static final long serialVersionUID = 1049383242803653757L;
    /**
     * 订单号
     */
    private Long orderNo;
    /**
     * 司机编号
     */
    private Long driverNo;
    /**
     * 乘客编号
     */
    private Long customerNo;
    /**
     * 订单总费用（计费的最终费用，已经扣除了所有计费规则上的优惠）
     */
    private Long totalFee = 0L;
    /**
     * 起步费
     */
    private Long startFee = 0L;
    /**
     * 总里程
     */
    private Double travelKm = 0d;
    /**
     * 里程费
     */
    private Long travelKmFee = 0L;
    /**
     * 行程时间
     */
    private Integer travelMinute = 0;
    /**
     * 行程时长费
     */
    private Long travelMinuteFee= 0L;
    /**
     * 低速时间（目前已废弃）
     */
    private Integer lowMinute = 0;
    /**
     * 低速费（目前已废弃）
     */
    private Long lowMinuteFee = 0L;
    /**
     * 长途里程
     */
    private Double longKm = 0d;
    /**
     * 长途费
     */
    private Long longKmFee = 0L;
    /**
     * 超远长途里程（已废弃）
     */
    private Double longKm2 = 0d;
    /**
     * 超远长途费（已废弃）
     */
    private Long longKm2Fee = 0L;
    /**
     * 夜间时长
     */
    private Integer nightTime = 0;
    /**
     * 夜间里程
     */
    private Double nightKm = 0d;
    /**
     * 夜间起步费
     */
    private Long nightFix = 0L;
    /**
     * 夜间里程费
     */
    private Long nightFee = 0L;
    /**
     * 停车费
     */
    private Long parkFee = 0L;
    /**
     * 高速费
     */
    private Long hightSpeedFee = 0L;
    /**
     * 过桥费
     */
    private Long bridgeFee = 0L;
    /**
     * 额外费
     */
    private Long extraFee = 0L;
    /**
     * 其他费用
     */
    private Long otherFee = 0L;
    /**
     * 取消费（目前公务没有）
     */
    private Long cancelFee = 0L;
    /**
     * 折扣优惠费用
     */
    private Long companyDiscountAmount=0L;
    /**
     * 折扣类型
     */
    private Byte derateType;
    /**
     * 折扣比率
     */
    private Integer derateRate;
    /**
     * 折扣金额
     */
    private Long derateAmount;
    /**
     * 改价一口价（改费直接改总价的时候记录用）
     */
    private Long fixedPrice;
    /**
     * 改价备注
     */
    private String remark;
    /**
     * 改价账单编号   
     */
    private Long updateBillNo;
    /**
     * 叠加夜间里程的方法
     * @param km
     */
    public void addNightKm(Double km){
        if(this.nightKm ==null){
            nightKm = 0.0;
        }
        this.nightKm += km;
    }

    /**
     * 订单费用（不包含司机费用的部分）
     * @return
     */
    @JSONField(serialize = false)
    @JsonIgnore
    public long getRestFee() {
        long restFee = 0;
        if(this.startFee != null) {
            restFee += startFee;
        }
        if(this.travelKmFee != null) {
            restFee += travelKmFee;
        }
        if(this.longKmFee != null) {
            restFee += longKmFee;
        }
        if(this.longKm2Fee != null) {
            restFee += longKm2Fee;
        }
        if(this.travelMinuteFee != null) {
            restFee += travelMinuteFee;
        }
        if(this.nightFee != null) {
            restFee += nightFee;
        }
        if(this.nightFix != null) {
            restFee += nightFix;
        }
        return restFee;
    }

    @JSONField(serialize = false)
    @JsonIgnore
    public long getLowestCostFee(){
        //按照目前的折扣金额规则，如果存在折扣减免，则预约单最低消费为0
        if(derateAmount == null || derateAmount>0){
            return 0;
        }
        long restFee = getRestFee();
        long lowestCost = this.totalFee - restFee -getDriverFee();
        if(lowestCost<0) {
            lowestCost = 0;
        }
        return lowestCost;
    }

    /**
     * 获取订单总费用（原价）
     * @return
     */
    @JSONField(serialize = false)
    @JsonIgnore
    public long getOriginalPrice() {
        return getRestFee()+getDriverFee();
    }

    /**
     * 获取司机费用
     * @return
     */
    @JSONField(serialize = false)
    @JsonIgnore
    public long getDriverFee(){
        long driverFee = 0;
        if(this.parkFee != null) {
            driverFee += parkFee;
        }
        if(this.hightSpeedFee != null) {
            driverFee += hightSpeedFee;
        }
        if(this.bridgeFee != null) {
            driverFee += bridgeFee;
        }
        if(this.extraFee != null) {
            driverFee += extraFee;
        }
        if(this.otherFee != null) {
            driverFee += otherFee;
        }
        return driverFee;
    }

}
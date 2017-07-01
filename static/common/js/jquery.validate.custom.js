//jquery.validation is custom validated
$(function(){
    //validator chinaese
    jQuery.validator.addMethod("chinaese", function(value, element) {
        var chinaese = /^[\u4e00-\u9fa5]+$/;
        return this.optional(element) || (chinaese.test(value));
    }, "请输入中文");
    //validator english
    jQuery.validator.addMethod("english", function(value, element) {
        var english = /^[A-Za-z]+$/;
        return this.optional(element) || (english.test(value));
    }, "请输入英文字符串");
});

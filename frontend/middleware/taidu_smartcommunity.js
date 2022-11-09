/*********************************************************
 * Taidu Technology Inc. © 2021
 * created at 12/18/2021
 * author: Yuwei Le
 * description: js functionalities for smart community project
 * dependencies: taidu_common.js
 *********************************************************/

/*
    CONST STRINGS
 */
const default_dep_jquery = "../middleware/jq-3.6.0.js";
const default_dep_crypto = "../middleware/crypto-js-4.1.1/crypto-js.js";

const desp_permission_super = "全域管理员";
const desp_permission_normal = "普通管理员";

const warn_info_not_complete = "您的信息未填写完整，请检查后重试！";
const warn_login_failed = "登录失败，请检查后重试！";
const warn_register_failed = "新用户注册失败，请检查后重试！";
const warn_username_too_short = "您输入的用户名太短，不可少于6位字符";
const warn_password_too_short = "您输入的密码太短，不可少于8位字符";
const warn_passwords_not_identical = "您两次输入的密码不符，请检查！";
const warn_already_login = "您已登录，即将为您跳转主页";
const warn_login_expired = "您的登录已过期，请重新登录！";

const text_version = "当前版本 V0.1<br>后端版本 V1.0.1";

const template_log_item = "<tr>\n" +
    "                    <th scope=\"row\">{0}</th>\n" +
    "                    <td>{1}</td>\n" +
    "                    <td>{2}</td>\n" +
    "                    <td>{3}</td>\n" +
    "                    <td>{4}</td>\n" +
    "                    <td>{5}</td>\n" +
    "                    <td>{6}</td>\n" +
    "                    <td>{7}</td>\n" +
    "                    <td>{8}</td>\n" +
    "                </tr>";


/**
 * 在指定区域显示提醒文字
 * @param warnTextArea 指定显示区域，是一个DOM对象
 * @param warnContent 显示提醒内容
 */
function showWarnText(warnTextArea, warnContent) {
    warnTextArea.style = "display: block";
    warnTextArea.innerText = warnContent;
}


/**
 * 检查一个用户token是否是合法的，如合法则返回关联用户信息
 * @param dep_jquery 如为null则用默认值，jQuery依赖包的相对路径
 * @param token 用户token
 * @param frontend_response function(code,data) 前端响应函数，实现code和data两个输入参数
 * @return {null|string} token的验证结果
 * @template
 * frontend_response = function (code, data) {
 *     if (isSuccess(code)) {
 *          alert("success");
 *          return data;
 *      } else {
 *          clearToken();
 *          return null;
 *      }
 * }
 */
function validateToken(dep_jquery, token, frontend_response) {
    if (dep_jquery === null)
        dep_jquery = default_dep_jquery;
    envPrepare(dep_jquery, function () {
        apiRequest({token: token}, const_api_list.postToken, frontend_response);
    });
}


/**
 * 根据token返回的权限id返回权限描述短语
 * @param permission_id token请求返回的权限id
 * @returns {string} 权限描述短语
 */
function getPermissionByID(permission_id) {
    switch(permission_id){
        case 0:
            return desp_permission_normal;
        case 1:
            return desp_permission_normal;
        case 2:
            return desp_permission_super;
    }
    return "";
}


/**
 * 登录业务后端函数
 * @param dep_jquery 如为null则用默认值，jQuery依赖包的相对路径
 * @param dep_crypto 如为null则用默认值，CryptoJS依赖包的相对路径
 * @param username 用户名
 * @param password 密码明文
 * @param warnTextArea 提醒区域，是一个DOM对象
 * @param urlAfterSuccess 成功登录后的自动跳转url
 */
function commandLogin(dep_jquery, dep_crypto, username, password, warnTextArea, urlAfterSuccess) {
    if (dep_jquery === null)
        dep_jquery = default_dep_jquery;
    if (dep_crypto === null)
        dep_crypto = default_dep_crypto;
    envPrepare(dep_jquery, function () {
        envPrepare(dep_crypto, function () {
            if (username === "" || password === "") {
                showWarnText(warnTextArea, warn_info_not_complete);
                return;
            }
            password = CryptoJS.MD5(password);
            apiRequest({username: username, password: password},
                const_api_list.postLogin,
                function (code, data){
                    if (isSuccess(code)) {
                        document.location.href = urlAfterSuccess;
                    } else {
                        showWarnText(warnTextArea, warn_login_failed);
                        showWarnText(warnTextArea, code + data);
                    }
                });
        });
    });
}


/**
 * 新用户注册业务后端函数
 * @param dep_jquery 如为null则用默认值，jQuery依赖包的相对路径
 * @param dep_crypto 如为null则用默认值，cryptoJS依赖包的相对路径
 * @param username 用户名
 * @param password 密码明文
 * @param warnTextArea 错误提醒文字区域，是一个DOM对象
 * @param urlAfterSuccess 成功注册后的跳转url(一般为登录页)
 * @returns {boolean} 是否成功注册
 */
function commandRegister(dep_jquery, dep_crypto, username, password, warnTextArea, urlAfterSuccess) {
    if (dep_jquery === null)
        dep_jquery = default_dep_jquery;
    if (dep_crypto === null)
        dep_crypto = default_dep_crypto;
    envPrepare(dep_jquery, function () {
       envPrepare(dep_crypto, function () {
           if (username === "" || password === "") {
               showWarnText(warnTextArea, warn_info_not_complete);
               return false;
           }
           if (username.length < 6) {
               showWarnText(warnTextArea, warn_username_too_short);
               return false;
           }
           if (password.length < 8) {
               showWarnText(warnTextArea, warn_password_too_short);
               return false;
           }
           password = CryptoJS.MD5(password);
           apiRequest({username: username, password: password},
               const_api_list.postAccount,
               function (code, data){
                   if (isSuccess(code)) {
                       document.location.href = urlAfterSuccess;
                       return true;
                   } else {
                       showWarnText(warnTextArea, warn_register_failed);
                       return false;
                   }
               });
       });
    });
}


/**
 * 按需获取日志记录并显示在tbody表格里
 * @param dep_jquery 如为null则用默认值，jQuery依赖包的相对路径
 * @param params 申请获取日志的参数
 * @param tbodyArea tbody标签，是一个DOM对象
 */
function commandLogSearch(dep_jquery, params, tbodyArea) {
    if (dep_jquery === null)
        dep_jquery = default_dep_jquery;
    envPrepare(dep_jquery, function () {
        apiRequest(params, const_api_list.getSearchLog, function (code, data) {
            if (isSuccess(code)) {
                tbodyArea.innerHTML = "";
                for (var i = 0; i < data.length; i++) {
                    tbodyArea.innerHTML += template_log_item.format(i + 1,
                        data[i].blocknumber,
                        data[i].uploadnode,
                        data[i].status,
                        data[i].lognumber,
                        data[i].date,
                        data[i].business,
                        data[i].content,
                        data[i].operator);
                }
            } else {
                tbodyArea.innerHTML += template_log_item.format(0, "-", "-", "-", "-", "-", "-", "-", "-");
            }
        });
    });
}


/**
 * API接口定时刷新
 * @param dep_jquery 如为null则用默认值，jQuery依赖包的相对路径
 * @param api 从const_api_list选择的定时刷新API
 * @param params API请求的随附参数
 * @param interval 刷新间隔（毫秒），不建议太短
 * @param textArea 显示区域，是一个DOM对象
 * @param callback (可为空)如有需要进一步解析返回值，则给好回调函数
 */
function refreshApiTask (dep_jquery, api, params, interval, textArea, callback) {
    if (dep_jquery === null)
        dep_jquery = default_dep_jquery;
    envPrepare(dep_jquery, function () {
        // run once
        apiRequest(params, api, function (code, data) {
            if (isSuccess(code)) {
                if (callback === null)
                    textArea.innerText = data;
                else
                    callback(data);
            }else {
                textArea.innerText = "-1";
            }
        });
        // set a repeating timer
        setInterval(function () {
            apiRequest(params, api, function (code, data) {
                if (isSuccess(code)) {
                    if (callback === null)
                        textArea.innerText = data;
                    else
                        callback(data);
                }else {
                    textArea.innerText = "-1";
                }
            })
        }, interval);
    });
}


/**
 * 获取当前时间字符串
 * @returns {string} 时间字符串
 */
function getFormattedTime() {
    var now = new Date();
    var year = now.getFullYear();
    var month = now.getMonth()+1;
    if (month < 10)
        month = "0" + month;
    var date = now.getDate();
    if (date < 10)
        date = "0" + date;
    var hour = now.getHours();
    if (hour < 10)
        hour = "0" + hour;
    var minute = now.getMinutes();
    if (minute<10)
        minute = "0" + minute;
    var second = now.getSeconds();
    if (second < 10)
        second="0"+second;
    return year + "-" + month + "-" + date + "<br>" + hour + ":" + minute + ":" + second;
}
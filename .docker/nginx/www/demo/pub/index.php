<?php
/**
 * 2012-2023 Group
 *
 * NOTICE OF LICENSE
 *
 * This source file is subject to the Open Software License (OSL 3.0) that is available
 * through the world-wide-web at this URL: http://www.opensource.org/licenses/OSL-3.0
 * If you are unable to obtain it through the world-wide-web, please send an email
 * so we can send you a copy immediately.
 *
 * @author fusc
 * @copyright 2012-2023
 * @license http://www.opensource.org/licenses/OSL-3.0 Open Software License (OSL 3.0)
 */

echo '<div class="center">';
// FT fonts support Check
// https://github.com/magento/magento2/issues/18177
if (!function_exists("imageftbbox")) {
    echo "<p style='width:934px;color:red;margin: 1em auto;'>Image CAPTCHA requires FT fonts support</p>" ;
} else {
    echo "<p style='width:934px;margin: 1em auto;'>Great, imageftbbox is exists.</p>";
}

// ICU Check
$locale = 'zh_Hans_CN';
$timeZone = 'Asia/Shanghai';
$dateStyle = $timeStyle = 3;
$formatter = new \IntlDateFormatter(
    $locale,
    $dateStyle,
    $timeStyle,
    $timeZone
);
$longYearPattern = preg_replace(
    '/(?<!y)yy(?!y)/',
    'y',
    (string)$formatter->getPattern()
);
$formatter->setPattern($longYearPattern);
$time = $formatter->parse('2023-06-16 00:00:00');
if ($time) {
    echo "<p style='width:934px;color:red;margin: 1em auto;'>ICU has problem.</p>";
    echo "<p style='width:934px;color:red;margin: 1em auto;'>ICU parse time: $time</p>";
    echo "<p style='width:934px;color:red;margin: 1em auto;'>date time: " . date('Y-m-d H:i:s', $time) . "</p>";
} else {
    echo "<p style='width:934px;margin: 1em auto;'> ICU is OK.</p>";
}
$locales = \ResourceBundle::getLocales('') ?: [];
echo '<p> Locales : '. count($locales) . '</p>';
echo '<p>';
foreach ($locales as $locale ) {
     echo $locale.'; ';
}
echo '</p>';
echo '</div>';

//phpinfo
phpinfo();

